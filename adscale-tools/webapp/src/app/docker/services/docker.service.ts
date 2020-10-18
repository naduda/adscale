import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { IDockerState } from '../model/state.interface';

@Injectable({
  providedIn: 'root'
})
export class DockerService {

  constructor(private http: HttpClient) { }

  get state$(): Observable<IDockerState> {
    return this.http.get<IDockerState>(`${environment.apiPreffix}/docker-state`);
  }

  toggleContainer(status: boolean): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/toggle-container`, { status });
  }

  createOrRemoveContainer(status: boolean): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/create-remove-container`, { status });
  }

  createOrRemoveImage(status: boolean): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/create-remove-image`, { status });
  }

  buildWar(installModule: boolean, installCbfsms: boolean): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/build-war`, { installModule, installCbfsms });
  }

  updateFrontend(isDev: boolean, installNpm: boolean): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/update-frontend`, { isDev, installNpm });
  }
}
