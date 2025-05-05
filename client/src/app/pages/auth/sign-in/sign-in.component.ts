import { ChangeDetectionStrategy, Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { ButtonDirective, InputDirective } from '@vm/ui';

@Component({
  selector: 'app-sign-in',
  imports: [InputDirective, ButtonDirective, RouterLink],
  templateUrl: './sign-in.component.html',
  styleUrl: './sign-in.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignInComponent {}
