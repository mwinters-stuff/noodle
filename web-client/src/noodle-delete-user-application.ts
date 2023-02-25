import { html, LitElement } from 'lit';
import { customElement, property, query, state } from 'lit/decorators.js';
import { Dialog } from '@material/mwc-dialog';
import '@material/mwc-button';
import { consume } from '@lit-labs/context';
import { noodleApiContext } from './noodle-context.js';
import { NoodleApiApi, ResponseError, UserApplications } from './api/index.js';

@customElement('noodle-delete-user-application')
export class NoodleDeleteUserApplication extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @query('mwc-dialog')
  private _dialog!: Dialog;

  @property()
  userApplication: UserApplications = {};

  @state()
  errorText: string = '';

  @state()
  appName: string = '';

  render() {
    return html`
      <mwc-dialog heading="Confirm Delete" @closed=${this.handleClosed}>
        <p>Are you sure you want to delete ${this.appName}</p>
        <p>${this.errorText}</p>
        <mwc-button slot="secondaryAction" dialogAction="close"
          >Cancel</mwc-button
        >
        <mwc-button slot="primaryAction" dialogAction="delete"
          >Delete</mwc-button
        >
      </mwc-dialog>
    `;
  }

  public show(userApp: UserApplications) {
    this.userApplication = userApp;
    this.appName = this.userApplication.Application?.Name || '';
    this.errorText = '';
    this._dialog.show();
  }

  private handleClosed(event: CustomEvent) {
    if (event.detail.action === 'delete') {
      this.errorText = '';
      this.noodleApi
        .noodleApplicationsDelete(this.userApplication.ApplicationId!)
        .then(() => {
          this.dispatchEvent(
            new Event('delete-user-application-dialog-closed')
          );
          this._dialog.close();
        })
        .catch(reason => {
          this.showError(reason);
        });
    }
  }

  private showError(reason: ResponseError) {
    reason.response.json().then((value: any) => {
      this.errorText = value.message;
    });
  }
}
