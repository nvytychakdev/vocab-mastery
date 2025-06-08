import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { Router } from '@angular/router';
import { Button } from '@vm/ui';
import { AuthService } from '../../../core/auth/auth.service';
import { FullscreenLayoutComponent } from '../../../layouts/fullscreen-layout/fullscreen-layout.component';

@Component({
  selector: 'app-confirm-email',
  imports: [FullscreenLayoutComponent, Button],
  templateUrl: './confirm-email.component.html',
  styleUrl: './confirm-email.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfirmEmailComponent implements OnInit {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);
  readonly email = signal<string | undefined>(undefined);

  ngOnInit(): void {
    const state = (this.router.getCurrentNavigation()?.extras?.state || window.history.state) as { email: string };
    if (state.email) this.email.set(state.email);
  }

  resendEmail() {
    const email = this.email();
    if (!email) return;
    this.auth.resendConfirmEmail(email).subscribe(console.log);
  }
}
