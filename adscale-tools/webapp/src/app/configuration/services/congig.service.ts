import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map, shareReplay, switchMap, tap } from 'rxjs/operators';
import { environment } from '../../../environments/environment';
import { IConfigurationProperty, ISettings } from '../components/model/interfaces';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  props$: Observable<IConfigurationProperty[]>;

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

  getProperties(forse = false): Observable<IConfigurationProperty[]> {
    if (!this.props$ || forse) {
      this.props$ = this.http.get<any>(`${environment.apiPreffix}/properties`)
        .pipe(
          map(e => this.mapProperties(e)),
          shareReplay(1),
        );
    }

    return this.props$;
  }

  private mapProperties(originalList: any): IConfigurationProperty[] {
    const res = Object.entries(originalList)
      .map((e: any[]) => ({ ...e[1], name: e[0] }));

    res.forEach(v => {
      const strVal = v.value as string;
      if (strVal.toLowerCase() === 'true' || strVal.toLowerCase() === 'false') {
        v.type = 'boolean';
        v.value = strVal.toLowerCase() === 'true';
      }
    });

    return res.sort((a, b) => {
      if (a.status && !b.status) {
        return -1;
      }
      if (!a.status && b.status) {
        return 1;
      }
      if (a.enabled && !b.enabled) {
        return -1;
      }
      if (!a.enabled && b.enabled) {
        return 1;
      }
      return a.name.localeCompare(b.name);
    });
  }

  saveProperties(data): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/properties`, data)
      .pipe(
        tap(_ => this.props$ = null),
      );
  }

  addProperty(property: { name: string, value: string }): Observable<IConfigurationProperty[]> {
    return this.http.post<void>(`${environment.apiPreffix}/add-property`, property)
      .pipe(
        switchMap(_ => this.getProperties(true)),
      );
  }

  removeLine(line: number): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/remove-property`, { line });
  }

  removeExtraLines(): Observable<IConfigurationProperty[]> {
    return this.http.post<any>(`${environment.apiPreffix}/remove-extra-empty-lines`, null)
      .pipe(
        map(e => this.mapProperties(e)),
      );
  }

  copyToContainer(): Observable<void> {
    return this.http.post<void>(`${environment.apiPreffix}/copy-properties-to-container`, null);
  }
}
