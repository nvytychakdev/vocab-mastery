import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  ValidationErrors,
  Validators,
} from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { ButtonDirective, InputDirective } from '@vm/ui';
import { AuthService, isConfirmResponse } from '../../../core/auth/auth.service';
import { FullscreenLayoutComponent } from '../../../layouts/fullscreen-layout/fullscreen-layout.component';

const passwordsMatch = (group: AbstractControl): ValidationErrors | null => {
  if (group instanceof FormGroup) {
    const isMatchingPwd = group.controls['password'].value !== group.controls['passwordRepeat'].value;
    return isMatchingPwd ? { passwordsMatch: true } : null;
  }
  return null;
};

@Component({
  selector: 'app-sign-up',
  imports: [ButtonDirective, InputDirective, RouterLink, FullscreenLayoutComponent, ReactiveFormsModule],
  templateUrl: './sign-up.component.html',
  styleUrl: './sign-up.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignUpComponent {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);

  readonly form = new FormGroup(
    {
      email: new FormControl('', [Validators.required]),
      name: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
      passwordRepeat: new FormControl('', [Validators.required]),
    },
    {
      validators: [passwordsMatch],
    }
  );

  submit() {
    const email = this.form.value.email;
    const name = this.form.value.name;
    const password = this.form.value.password;
    const passwordRepeat = this.form.value.passwordRepeat;

    if (!email || !password || !name) throw new Error('No password, name or email provided');
    if (passwordRepeat !== password) throw new Error('Password does not match');

    this.auth.signUp({ email, password, name }).subscribe(res => {
      if (isConfirmResponse(res)) {
        void this.router.navigate(['/auth/confirm-email'], { state: { email } });
        return;
      }

      void this.router.navigate(['/auth/sign-in']);
    });
  }
}
