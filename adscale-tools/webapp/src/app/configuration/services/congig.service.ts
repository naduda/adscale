import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from '../../../environments/environment';
import { IConfigurationProperty, ISettings } from '../components/model/interfaces';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  constructor(private http: HttpClient) { }

  autoCompleteFilePath(path: string): Observable<string[]> {
    return this.http.post<string[]>(`${environment.apiPreffix}/file-path-autocomplete`, { path });
  }

  getSettings(): Observable<ISettings> {
    return this.http.get<ISettings>(`${environment.apiPreffix}/settings`);
  }

  saveSettings(data: ISettings): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/settings`, data);
  }

  getProperties(): Observable<IConfigurationProperty[]> {
    return this.http.get<any>(`${environment.apiPreffix}/properties`)
      .pipe(
        map(e => {
          const res: IConfigurationProperty[] = [];
          Object.keys(e).forEach(k => {
            const v: IConfigurationProperty = e[k];
            const strVal = v.value as string;
            if (strVal.toLowerCase() === 'true' || strVal.toLowerCase() === 'false') {
              v.type = 'boolean';
              v.value = strVal.toLowerCase() === 'true';
            }
            res.push({ ...v, name: k });
          });
          return res.sort((a, b) => a.name.localeCompare(b.name));
        }),
      );
  }

  saveProperties(data): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/properties`, data);
  }
}
