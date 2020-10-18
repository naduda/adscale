import { Component, OnInit } from '@angular/core';
import { finalize } from 'rxjs/operators';
import { IDockerState } from '../../model/state.interface';
import { DockerService } from '../../services/docker.service';

@Component({
  selector: 'adscale-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.sass']
})
export class CreateComponent implements OnInit {

  state: IDockerState;
  loading: boolean;

  installModule: boolean;
  installCbfsms: boolean;
  installNpm: boolean;

  constructor(
    private dockerService: DockerService,
  ) { }

  ngOnInit(): void {
    this.dockerService.state$
      .subscribe(e => this.state = e);
  }

  toggleContainer(v: boolean) {
    this.dockerService.toggleContainer(v).subscribe(_ => {
      this.state.containerRunning = !this.state.containerRunning;
    });
  }

  createOrRemoveContainer() {
    this.loading = true;
    this.dockerService.createOrRemoveContainer(!this.state.containerExists)
      .pipe(finalize(() => this.loading = false))
      .subscribe(_ => {
        this.state.containerExists = !this.state.containerExists;
        this.state.containerRunning = this.state.containerExists;
      });
  }

  createOrRemoveImage() {
    this.loading = true;
    this.dockerService.createOrRemoveImage(!this.state.imageExists)
      .pipe(finalize(() => this.loading = false))
      .subscribe(_ => {
        this.state.imageExists = !this.state.imageExists;
        this.state.containerExists = this.state.imageExists;
        this.state.containerRunning = this.state.containerExists;
      });
  }

  buildWar() {
    this.loading = true;
    this.dockerService.buildWar(this.installModule, this.installCbfsms)
      .pipe(finalize(() => this.loading = false))
      .subscribe();
  }

  updateFrontend(isDev: boolean) {
    this.loading = true;
    this.dockerService.updateFrontend(isDev, this.installNpm)
      .pipe(finalize(() => this.loading = false))
      .subscribe();
  }

}
