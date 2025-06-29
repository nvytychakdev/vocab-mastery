import { ChangeDetectionStrategy, Component, DestroyRef, inject } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { AuthService, isConfirmResponse } from '@feature/auth/auth.service';
import { Button, Input } from '@vm/ui';
import { FullscreenLayoutComponent } from '../../../layouts/fullscreen-layout/fullscreen-layout.component';

@Component({
  selector: 'app-sign-in',
  imports: [Input, Button, RouterLink, FullscreenLayoutComponent, ReactiveFormsModule],
  templateUrl: './sign-in.component.html',
  styleUrl: './sign-in.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignInComponent {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);
  private readonly destroyRef = inject(DestroyRef);

  readonly form = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  submit() {
    const email = this.form.value.email;
    const password = this.form.value.password;

    if (!email || !password) throw new Error('No password or email provided');

    this.auth.signIn(email, password).subscribe(res => {
      if (isConfirmResponse(res)) {
        void this.router.navigate(['/auth/confirm-email'], { state: { email } });
        return;
      }
      void this.router.navigate(['/main']);
    });
  }

  signInWithGoogle() {
    this.auth
      .signInWithGoogle()
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(() => {
        void this.router.navigate(['/main']);
      });
  }
}
