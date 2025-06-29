import { ChangeDetectionStrategy, Component, computed, inject } from '@angular/core';
import { Router } from '@angular/router';
import { AuthProfileState } from '@domain/auth/auth-profile-state';
import { AuthService } from '@feature/auth/auth.service';
import { NgIcon } from '@ng-icons/core';
import { Menu, MenuItem, MenuTrigger } from '@vm/ui';

@Component({
  selector: 'app-main-layout-sidebar-profile',
  imports: [NgIcon, MenuTrigger, Menu, MenuItem],
  templateUrl: './main-layout-sidebar-profile.component.html',
  styleUrl: './main-layout-sidebar-profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MainLayoutSidebarProfileComponent {
  private readonly auth = inject(AuthService);
  private readonly authProfile = inject(AuthProfileState);
  private readonly router = inject(Router);

  readonly profileName = computed(() => this.authProfile.profile()?.name);
  readonly profileEmail = computed(() => this.authProfile.profile()?.email);
  readonly profilePictureUrl = computed(() => this.authProfile.profile()?.pictureUrl);

  signOut() {
    this.auth.signOut().subscribe(() => this.router.navigate(['/']));
  }
}
