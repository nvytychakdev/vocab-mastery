import { Overlay, OverlayRef, STANDARD_DROPDOWN_BELOW_POSITIONS } from '@angular/cdk/overlay';
import { TemplatePortal } from '@angular/cdk/portal';
import {
  DestroyRef,
  Directive,
  ElementRef,
  HostListener,
  inject,
  input,
  OnInit,
  ViewContainerRef,
} from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Menu } from './menu';

@Directive({
  selector: '[vmMenuTrigger]',
})
export class MenuTrigger implements OnInit {
  private readonly overlay = inject(Overlay);
  private readonly viewContainerRef = inject(ViewContainerRef);
  private readonly elementRef = inject<ElementRef<HTMLElement>>(ElementRef);
  private readonly destroyRef = inject(DestroyRef);

  readonly menu = input.required<Menu>({ alias: 'vmMenuTrigger' });

  private overlayRef: OverlayRef | null = null;

  ngOnInit() {
    this.menu()
      .open.pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(() => this.open());
  }

  @HostListener('click')
  onClick() {
    this.open();
  }

  open() {
    if (!this.overlayRef) {
      this.overlayRef = this.createOverlay();
    }

    if (this.overlayRef.hasAttached()) return;

    this.attachMenu();
    this.menu().isOpen.set(true);

    this.menu()
      .close.pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(() => this.close());

    this.overlayRef
      ?.outsidePointerEvents()
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(event => {
        if (event.target === this.elementRef.nativeElement) return;
        if (event.target && this.elementRef.nativeElement.contains(event.target as Element)) return;
        this.close();
      });
  }

  close() {
    if (!this.overlayRef) return;
    this.menu().isOpen.set(false);
    this.overlayRef.detach();
    this.overlayRef = null;
  }

  private createOverlay() {
    const overlayRef = this.overlay.create({
      scrollStrategy: this.overlay.scrollStrategies.reposition(),
      positionStrategy: this.overlay
        .position()
        .flexibleConnectedTo(this.elementRef)
        .withFlexibleDimensions(false)
        .withLockedPosition()
        .withPositions(STANDARD_DROPDOWN_BELOW_POSITIONS),
    });

    return overlayRef;
  }

  private attachMenu() {
    if (!this.overlayRef) return;
    const portal = new TemplatePortal(this.menu().templateRef(), this.viewContainerRef);
    this.overlayRef.attach(portal);
  }
}
