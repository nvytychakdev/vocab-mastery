import { Directive, HostBinding, input } from '@angular/core';

type ButtonSize = 'small' | 'medium' | 'large';
type ButtonType = 'button' | 'link' | 'icon';
type ButtonVariant = 'primary' | 'secondary' | 'warn' | 'danger' | 'ghost';

const ButtonVariants: Record<ButtonVariant, string> = {
  primary: 'vm-button-primary',
  secondary: 'vm-button-secondary',
  danger: 'vm-button-danger',
  warn: 'vm-button-warn',
  ghost: 'vm-button-ghost',
};

const ButtonSizes: Record<ButtonSize, string> = {
  small: 'vm-button-sm',
  medium: 'vm-button-md',
  large: 'vm-button-lg',
};

const ButtonTypes: Record<ButtonType, string> = {
  button: 'vm-button',
  link: 'vm-button-link',
  icon: 'vm-button vm-button-icon',
};

@Directive({
  selector: '[vmButton]',
})
export class ButtonDirective {
  readonly size = input<ButtonSize>('large');
  readonly type = input<ButtonType>('button');
  readonly variant = input<ButtonVariant>('secondary');

  @HostBinding('class') get className() {
    if (this.type() === 'link') {
      return `${ButtonTypes[this.type()]} ${ButtonSizes[this.size()]}`;
    }
    return `${ButtonTypes[this.type()]} ${ButtonVariants[this.variant()]} ${ButtonSizes[this.size()]}`;
  }
}
