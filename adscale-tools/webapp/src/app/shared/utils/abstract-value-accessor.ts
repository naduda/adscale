import { ControlValueAccessor } from '@angular/forms';

export abstract class AValueAccessor implements ControlValueAccessor {

  onChange = (_: any) => { };
  onTouched = () => { };
  setDisabledState?(isDisabled: boolean): void;
  registerOnChange(fn: (_: any) => void): void { this.onChange = fn; }
  registerOnTouched(fn: () => void): void { this.onTouched = fn; }

  abstract writeValue(value: any): void;
  abstract onUpdate(value?: any): void;
}
