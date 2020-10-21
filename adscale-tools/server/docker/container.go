package docker

import (
	"adscale-tools/fileutils"
	"adscale-tools/model"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type AdscaleContainer struct {
	State model.DockerState
}

const imageTemplate = `
FROM tomcat:8.5
MAINTAINER adscale

ADD configuration /adscale/configuration

RUN keytool -genkey -alias adscale \
    -keyalg RSA -keystore /adscale/keystore.jks \
    -dname "CN=Adscale IL, OU=JavaSoft, O=Sun, L=Cupertino, S=Ukraine, C=UA" \
    -storepass 12345678 -keypass 12345678 \
    -ext SAN=dns:localhost.adscale.com

RUN sed -i 's|<Service name="Catalina">|<Service name="Catalina">\n\n<Connector SSLEnabled="true" acceptCount="100" clientAuth="false" disableUploadTimeout="true" enableLookups="false" maxThreads="25" port="8443" keystoreFile="/adscale/keystore.jks" keystorePass="12345678" protocol="org.apache.coyote.http11.Http11NioProtocol" scheme="https" secure="true" sslProtocol="TLS" />|g' /usr/local/tomcat/conf/server.xml

`

func (c *AdscaleContainer) Init() error {
	imageExists, err := isImageExists()
	if err != nil {
		return err
	}
	containerExists, err := isContainerExists()
	if err != nil {
		return err
	}
	containerRunning, err := isContainerRunning()
	if err != nil {
		return err
	}

	c.State = model.DockerState{
		ImageExists:      imageExists,
		ContainerExists:  containerExists,
		ContainerRunning: containerRunning,
	}

	return nil
}

func (c *AdscaleContainer) CreateImage(s model.Settings) error {
	if c.State.ImageExists {
		return nil
	}

	if err := fileutils.MakeDirIfNotExist("./" + model.DockerFolder); err != nil {
		return err
	}

	if err := PrepareEasyleadsConf(s); err != nil {
		return err
	}

	dockerfile := fmt.Sprintf("./%s/%s", model.DockerFolder, "Dockerfile")
	file, err := os.OpenFile(dockerfile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(imageTemplate)
	if err != nil {
		fmt.Println(err)
		return err
	}

	command := fmt.Sprintf("docker image build -t %s .", model.DockerImageName)
	if err = RunCommand(command, model.DockerFolder); err == nil {
		c.State.ImageExists = true
	}

	return err
}

func (c *AdscaleContainer) RemoveImage() error {
	command := fmt.Sprintf("docker rmi %s && docker image prune -f", model.DockerImageName)
	_, err := executeCommand(command)
	if err == nil {
		c.State.ImageExists = false
	}
	return err
}

func isImageExists() (bool, error) {
	res, err := executeCommand("docker images --format {{.Repository}}")
	if err != nil {
		return false, err
	}
	names := strings.Split(res, "\n")
	for _, n := range names {
		if n == model.DockerImageName {
			return true, nil
		}
	}
	return false, nil
}

func isContainerExists() (bool, error) {
	res, err := executeCommand("docker ps -a --format {{.Names}}")
	if err != nil {
		return false, err
	}
	names := strings.Split(res, "\n")
	for _, n := range names {
		if n == model.DockerContainerName {
			return true, nil
		}
	}
	return false, nil
}

func isContainerRunning() (bool, error) {
	res, err := executeCommand("docker ps --format {{.Names}}")
	if err != nil {
		return false, err
	}
	names := strings.Split(res, "\n")
	for _, n := range names {
		if n == model.DockerContainerName {
			return true, nil
		}
	}
	return false, nil
}

func (c *AdscaleContainer) CreateContainer(port int) error {
	appPort := strconv.Itoa(port)
	command := fmt.Sprintf("docker run --name %s -d --restart always --add-host %s:127.0.0.1 -p %s:8443 %s",
		model.DockerContainerName, model.DockerDomainUrl, appPort, model.DockerImageName)
	if err := RunCommand(command, "./"); err != nil {
		return err
	}

	c.State.ContainerExists = true
	return nil
}

func (c *AdscaleContainer) CreateAndRunContainer(port int) error {
	if err := c.CreateContainer(port); err != nil {
		return err
	}
	return c.ToggleContainer(true)
}

func (c *AdscaleContainer) StopAndRemoveContainer() error {
	if c.State.ContainerRunning {
		if err := c.ToggleContainer(false); err != nil {
			return err
		}
	}

	return c.RemoveContainer()
}

func (c *AdscaleContainer) ToggleContainer(state bool) error {
	dockerCommand := "stop"
	if state {
		dockerCommand = "start"
	}
	command := fmt.Sprintf("docker %s %s", dockerCommand, model.DockerContainerName)
	_, err := executeCommand(command)
	if err == nil {
		c.State.ContainerRunning = state
	}
	return err
}

func (c *AdscaleContainer) RemoveContainer() error {
	command := fmt.Sprintf("docker stop %s && docker rm %s && docker container prune -f",
		model.DockerContainerName, model.DockerContainerName)
	_, err := executeCommand(command)
	if err == nil {
		c.State.ContainerExists = false
	}
	return err
}

func findConnectionUrl(filename string) (string, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	str := string(input)
	idxBeg := strings.Index(str, "jdbc:mysql") + 13
	res := str[idxBeg:]
	idxEnd := strings.Index(str, ":3306")
	res = str[idxBeg:idxEnd]
	return res, nil
}

func CreateWar(mod3 bool, cbfsms bool, s model.Settings) error {
	contextXML := s.Repo + "/ui/WebContent/META-INF/context.xml"
	url, err := findConnectionUrl(contextXML)
	if err != nil {
		return err
	}
	if err := fileutils.ReplaceInFile(contextXML, url, s.DbIP); err != nil {
		return err
	}

	var command string
	if cbfsms {
		command += `mvn install:install-file \
-Dfile=/Users/slavapoet/projects/git.public/ui/src/main/resources/libs/cbfsms.jar \
-DgroupId=cbfsms -DartifactId=cbfsms -Dversion=1 -Dpackaging=jar && \`
	}

	if mod3 {
		command += `mvn install:install-file \
-Dfile=/Users/pr/projects/git.public/ui/src/main/resources/libs/adscale_modules-3.0.jar \
-DgroupId=com.adscale -DartifactId=adscale_modules -Dversion=3.0 -Dpackaging=jar && \`
	}

	renameCommand := fmt.Sprintf("mv %s/ui/target/adscale_ui-3.0.war %s/ui/target/ROOT.war",
		s.Repo, s.Repo)
	if runtime.GOOS == "windows" {
		renameCommand = fmt.Sprintf("ren %s/ui/target/adscale_ui-3.0.war ROOT.war", s.Repo)
		renameCommand = strings.Replace(renameCommand, "/", "\\", -1)
	}
	command += fmt.Sprintf(`mvn clean install -Dmaven.test.skip=true -f %s/base/pom.xml && \
mvn clean install -Dmaven.test.skip=true -f %s/ui/pom.xml`, s.Repo, s.Repo)

	command += fmt.Sprintf(` && \
%s && docker cp %s/ui/target/ROOT.war %s:/usr/local/tomcat/webapps/ROOT.war`,
		renameCommand, s.Repo, model.DockerContainerName)

	if runtime.GOOS == "windows" {
		command = strings.Replace(command, "\\\n", "", -1)
	}

	if err := RunCommand(command, s.Repo); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("war was updated inside container")

	return fileutils.ReplaceInFile(contextXML, s.DbIP, url)
}
