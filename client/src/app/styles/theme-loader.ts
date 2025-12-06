import { Injectable, provideAppInitializer, signal } from '@angular/core';
import { inject } from '@angular/core/primitives/di';

/**
 * Local storage theme key.
 */
const THEME_KEY = 'theme' as const;

/**
 * List of supported themes by the app.
 */
const SupportedThemes = { Light: 'light', Dark: 'dark' } as const;
const AllThemePreferences = { ...SupportedThemes, System: 'system' } as const;
type ThemeType = (typeof SupportedThemes)[keyof typeof SupportedThemes];
type ThemePreferencesType = (typeof AllThemePreferences)[keyof typeof AllThemePreferences];

@Injectable({ providedIn: 'root' })
export class ThemeLoader {
  private readonly _theme = signal<ThemeType | undefined>(undefined);
  readonly theme = this._theme.asReadonly();

  private readonly _preference = signal<ThemePreferencesType>('system');
  readonly preference = this._preference.asReadonly();

  cycleThemePreference() {
    if (this.preference() === AllThemePreferences.Dark) this.setThemePreference(AllThemePreferences.Light);
    else if (this.preference() === AllThemePreferences.Light) this.removeThemePreference();
    else this.setThemePreference(AllThemePreferences.Dark);
  }

  /**
   * Set explicit theme preference.
   * Theme will be stored in local stirage.
   * @param theme "dark" or "light"
   */
  setThemePreference(theme: ThemeType) {
    localStorage.setItem(THEME_KEY, theme);
    document.documentElement.classList.toggle(SupportedThemes.Dark, theme === SupportedThemes.Dark);
    this._theme.set(theme);
    this._preference.set(theme);
  }

  /**
   * Disable any theme explicit preferences.
   * The theme to use will be based on system preferences.
   */
  removeThemePreference() {
    localStorage.removeItem(THEME_KEY);
    const isQueryDarkTheme = window.matchMedia('(prefers-color-scheme: dark)').matches;
    document.documentElement.classList.toggle(SupportedThemes.Dark, isQueryDarkTheme);
    this._theme.set(isQueryDarkTheme ? SupportedThemes.Dark : SupportedThemes.Light);
    this._preference.set(AllThemePreferences.System);
  }
}

/**
 * Load theme on initialize of the app.
 */
export const loadTheme = () => {
  const themeLoader = inject(ThemeLoader);
  const localTheme = localStorage.getItem(THEME_KEY);
  const hasLocaltheme = localTheme && (localTheme === SupportedThemes.Dark || localTheme === SupportedThemes.Light);
  if (hasLocaltheme) themeLoader.setThemePreference(localTheme as ThemeType);
  else themeLoader.removeThemePreference();
};

/**
 * Provide angular's DI theme loader with the app.
 * Theme selection will be executed on app initizlization.
 * @returns app initializer func
 */
export const provideThemeLoader = () => provideAppInitializer(loadTheme);
