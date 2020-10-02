import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { Technology } from './tech.model';

@Injectable({
  providedIn: 'root'
})
export class TechService {

  constructor(private readonly httpClient: HttpClient) { }

  getTechnologies(): Observable<Technology[]> {
    return this.httpClient.get<Technology[]>(`${environment.apiPreffix}/technologies`);
  }
}
