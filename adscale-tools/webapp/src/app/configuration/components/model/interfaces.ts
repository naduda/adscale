export interface ISettings {
  easyleads: string;
  repo: string;
}

export interface IPath {
  path: string;
  name: string;
}

export interface IConfigurationProperty {
  line: number;
  name: string;
  value: string | boolean;
  status: boolean;
  enabled: boolean;
  type: string;
}
