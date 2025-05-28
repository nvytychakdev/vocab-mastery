import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { ComponentRef, inject, Injectable, Injector } from '@angular/core';
import { ToastCard } from './toast-card';
import { ToastContainer, ToastData } from './toast-container';

const TOAST_OFFSET_PX = '0px';

@Injectable({
  providedIn: 'root',
})
export class ToastService {
  private readonly injector = inject(Injector);
  private readonly overlay = inject(Overlay);

  private overlayRef?: OverlayRef;
  private containerRef?: ComponentRef<ToastContainer>;

  info(title: string, description?: string, duration?: number) {
    this.show({ title, description, duration });
  }

  success(title: string, description?: string, duration?: number) {
    this.show({ title, description, duration, type: 'success' });
  }

  warn(title: string, description?: string, duration?: number) {
    this.show({ title, description, duration, type: 'warn' });
  }

  error(title: string, description?: string, duration?: number) {
    this.show({ title, description, duration, type: 'error' });
  }

  private show(data: ToastData) {
    if (!this.overlayRef) {
      this.overlayRef = this.createOverlay();
      this.attachContainer();
    }

    this.containerRef?.instance.addToast(ToastCard, data);
  }

  private createOverlay() {
    const overlayRef = this.overlay.create({
      hasBackdrop: false,
      positionStrategy: this.overlay.position().global().bottom(TOAST_OFFSET_PX).right(TOAST_OFFSET_PX),
    });

    return overlayRef;
  }

  private attachContainer() {
    if (!this.overlayRef) return;
    const portal = new ComponentPortal(ToastContainer, null, this.injector);
    this.containerRef = this.overlayRef.attach(portal);
  }
}
