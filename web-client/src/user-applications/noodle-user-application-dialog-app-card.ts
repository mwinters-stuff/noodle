
import { customElement } from 'lit/decorators.js';
import { NoodleAppCard } from '../common/noodle-app-card.js';

@customElement('noodle-user-application-dialog-app-card')
export class NoodleUserApplicationDialogAppCard extends NoodleAppCard {

  public uploadIcon(file: File) {
    this.imageElement.src = URL.createObjectURL(file);
    this.imageElement.onload = () => {
      URL.revokeObjectURL(this.imageElement.src);
      this.imageElement.onload = null;
    };
  }
}

