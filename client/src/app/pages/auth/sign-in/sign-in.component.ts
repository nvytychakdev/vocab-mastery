import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { ButtonDirective, InputDirective } from '@vm/ui';
import { AuthService } from '../../../core/auth/auth.service';
import { FullscreenLayoutComponent } from '../../../layouts/fullscreen-layout/fullscreen-layout.component';

@Component({
  selector: 'app-sign-in',
  imports: [InputDirective, ButtonDirective, RouterLink, FullscreenLayoutComponent, ReactiveFormsModule],
  templateUrl: './sign-in.component.html',
  styleUrl: './sign-in.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignInComponent {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);

  readonly form = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  submit() {
    const email = this.form.value.email;
    const password = this.form.value.password;

    if (!email || !password) throw new Error('No password or email provided');

    this.auth.signIn(email, password).subscribe(() => {
      void this.router.navigate(['/main']);
    });
  }
}
