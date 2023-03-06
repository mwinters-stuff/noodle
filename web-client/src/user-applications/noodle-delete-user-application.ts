import { html, css, LitElement } from 'lit';
import { customElement, property, query, state } from 'lit/decorators.js';

import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/input/input.js';
import '@shoelace-style/shoelace/dist/components/dialog/dialog.js';
import '@shoelace-style/shoelace/dist/components/select/select.js';
import '@shoelace-style/shoelace/dist/components/option/option.js';

import SlDialog from '@shoelace-style/shoelace/dist/components/dialog/dialog';

import { consume } from '@lit-labs/context';
import { noodleApiContext } from '../noodle-context.js';
import { NoodleApiApi, ResponseError, UserApplications } from '../api/index.js';
import { Functions } from '../common/functions.js';

@customElement('noodle-delete-user-application')
export class NoodleDeleteUserApplication extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @query('sl-dialog')
  _dialog!: SlDialog;

  @property()
  userApplication: UserApplications = {};

  @state()
  errorText: string = '';

  @state()
  appName: string = '';

  static styles = css`
    .no-close::part(close-button) {
      visibility: hidden;
    }
  `
  render() {
    return html`
      <sl-dialog class="no-close" label="Confirm Delete" @sl-request-close=${this.dialogRequestClose}
        @sl-hide=${this.dialogClosed}>
        <p>Are you sure you want to delete ${this.appName}</p>
        <p>${this.errorText}</p>
        <sl-button slot="footer" variant="primary" @click=${this.primaryButtonClick}>Yes
        </sl-button>
        <sl-button slot="footer" variant="default" @click=${this.secondaryButtonClick}>No
        </sl-button>
      </sl-dialog>
    `;
  }

  public show(userApp: UserApplications) {
    this.userApplication = userApp;
    this.appName = this.userApplication.Application?.Name || '';
    this.errorText = '';
    this._dialog.show();
  }


  private dialogRequestClose(event: CustomEvent) {
    if (event.detail.source === 'overlay') {
      event.preventDefault();
    }
  }

  private dialogClosed(event: CustomEvent) {
    if (event.target == this._dialog) {
      this.errorText = '';
      this.dispatchEvent(new Event('delete-user-application-dialog-closed'));
    }
  }

  private secondaryButtonClick() {
    this._dialog.hide();
  }

  private primaryButtonClick() {
    this.errorText = '';
    this.noodleApi
      .noodleApplicationsDelete(this.userApplication.ApplicationId!)
      .then(() => {
        this.dispatchEvent(
          new Event('delete-user-application-dialog-closed')
        );
        this._dialog.hide();
      })
      .catch(reason => {
        Functions.showWebResponseError(reason);
      });
  }


}
