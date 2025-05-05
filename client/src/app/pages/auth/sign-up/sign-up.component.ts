import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonDirective, InputDirective } from '@vm/ui';

@Component({
  selector: 'app-sign-up',
  imports: [ButtonDirective, InputDirective, RouterLink],
  templateUrl: './sign-up.component.html',
  styleUrl: './sign-up.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignUpComponent {}
