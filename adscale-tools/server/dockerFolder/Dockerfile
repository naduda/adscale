
FROM tomcat:8.5
MAINTAINER adscale

ADD configuration /adscale/configuration

RUN keytool -genkey -alias adscale \
    -keyalg RSA -keystore /adscale/keystore.jks \
    -dname "CN=Adscale IL, OU=JavaSoft, O=Sun, L=Cupertino, S=Ukraine, C=UA" \
    -storepass 12345678 -keypass 12345678 \
    -ext SAN=dns:localhost.adscale.com

RUN sed -i 's|<Service name="Catalina">|<Service name="Catalina">\n\n<Connector SSLEnabled="true" acceptCount="100" clientAuth="false" disableUploadTimeout="true" enableLookups="false" maxThreads="25" port="8443" keystoreFile="/adscale/keystore.jks" keystorePass="12345678" protocol="org.apache.coyote.http11.Http11NioProtocol" scheme="https" secure="true" sslProtocol="TLS" />|g' /usr/local/tomcat/conf/server.xml

