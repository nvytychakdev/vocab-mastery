import {
  ChangeDetectionStrategy,
  Component,
  ComponentRef,
  signal,
  Type,
  viewChild,
  ViewContainerRef,
} from '@angular/core';
import { race, timer } from 'rxjs';
import { BaseToast, ToastType } from './toast-card';

export type ToastData = {
  title: string;
  description?: string;
  duration?: number;
  type?: ToastType;
};

const ToastDefaultDuration = 5000;
const ToastMaxCount = 5;

@Component({
  selector: 'vm-toast-container',
  imports: [],
  template: `
    <div class="vm-toast-container">
      <ng-container #container />
    </div>
  `,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ToastContainer {
  readonly containerRef = viewChild('container', { read: ViewContainerRef });
  readonly instances = signal<ComponentRef<BaseToast>[]>([]);

  addToast(component: Type<BaseToast>, data: ToastData) {
    const componentRef = this.containerRef()?.createComponent(component);
    if (!componentRef) throw new Error('Not able to create toast component');
    this.instances.update(i => [...i, componentRef]);

    // init
    componentRef.setInput('title', data.title);
    componentRef.setInput('description', data.description);
    componentRef.setInput('type', data.type);
    componentRef.instance.show();

    // remove extra instances from the list forcefully
    if (this.instances().length > ToastMaxCount) {
      const lastInstance = this.instances()
        .filter(({ instance }) => !instance.isRemoving)
        .at(0);
      if (lastInstance) this.removeToast(lastInstance);
    }

    // auto-removal of the instances
    const timer$ = timer(data?.duration ?? ToastDefaultDuration);
    race([timer$, componentRef.instance.onClose$]).subscribe(() => {
      if (!this.instances().includes(componentRef)) return;
      this.removeToast(componentRef);
    });
  }

  removeToast(componentRef: ComponentRef<BaseToast>) {
    componentRef.instance.hide();
    componentRef.instance.onRemove$.subscribe(() => {
      if (!this.containerRef()?.length) throw new Error('Container does not have any elements');
      const position = this.containerRef()?.indexOf(componentRef.hostView);
      this.containerRef()?.detach(position)?.destroy();
      this.instances.update(i => i.filter(ref => ref !== componentRef));
    });
  }
}
