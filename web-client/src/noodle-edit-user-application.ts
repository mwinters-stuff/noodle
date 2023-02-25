import { html, css, LitElement } from 'lit';
import { query, customElement, state, property } from 'lit/decorators.js';

import * as mwcButton from '@material/mwc-button';
import * as mwcDialog from '@material/mwc-dialog';
import * as mwcTextField from '@material/mwc-textfield';
import * as mwcSelect from '@material/mwc-select';

import { consume } from '@lit-labs/context';
import {
  DataCache,
  dataCacheContext,
  noodleApiContext,
  userSessionContext,
} from './noodle-context.js';
import {
  Application,
  ApplicationTemplate,
  NoodleApiApi,
  NoodleApplicationsPostActionEnum,
  NoodleApplicationTabsPostActionEnum,
  ResponseError,
  UserSession,
} from './api/index.js';

import '@material/mwc-dialog';
import '@material/mwc-button';
import '@material/mwc-textfield';
import '@material/mwc-formfield';
import '@material/mwc-select';

@customElement('noodle-edit-user-application')
export class NoodleEditUserApplication extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @consume({ context: userSessionContext, subscribe: true })
  @state()
  userSession!: UserSession;

  @consume({ context: dataCacheContext, subscribe: true })
  @state()
  dataCache!: DataCache;

  @state()
  _appTemplates!: ApplicationTemplate[];

  @state()
  _icon!: string;

  @state()
  selectedTabIndex: number = 0;

  @state()
  errorText: string = '';

  @property()
  application: Application = {};

  @query('#primary-action-button')
  _primaryButton!: mwcButton.Button;

  @query('#dialog')
  _dialog!: mwcDialog.Dialog;

  @query('#text-field-application-name')
  _textFieldApplicationName!: mwcTextField.TextField;

  @query('#text-field-application-url')
  _textFieldApplicationUrl!: mwcTextField.TextField;

  @query('#text-field-application-background')
  _textFieldBackground!: mwcTextField.TextField;

  @query('#text-field-application-icon')
  _textFieldIcon!: mwcTextField.TextField;

  @query('#select-tab')
  _selectTab!: mwcSelect.Select;

  static styles = css`
    div.vertical {
      display: flex;
      flex-direction: column;
    }
    mwc-textfield {
      margin-bottom: 16px;
      min-width: 600px;
    }
    mwc-select {
      margin-top: 16px;
      margin-bottom: 16px;
    }
    mwc-dialog {
      --mdc-dialog-max-width: 800px;
      --mdc-dialog-min-width: 800px;
    }
  `;

  public show(application: Application, tabId: number) {
    this.application = application;
    // this._textFieldApplicationName.textContent = application.name!;
    this._selectTab.select(this.dataCache.GetTabIndex(tabId));
    this._dialog.show();
  }

  private primaryButtonClick() {
    this.errorText = '';
    const isValid = this._textFieldApplicationName.checkValidity();
    if (isValid) {
      this.application.Name = this._textFieldApplicationName.value;
      this.application.Website = this._textFieldApplicationUrl.value;
      this.application.Description = this._textFieldApplicationName.value;
      this.application.TileBackground = this._textFieldBackground.value;
      this.application.Icon = this._textFieldIcon.value;
      if (this.application.TemplateAppid === '') {
        this.application.TemplateAppid = undefined;
      }

      this.noodleApi
        .noodleApplicationsPost(
          NoodleApplicationsPostActionEnum.Update,
          this.application
        )
        .then(appResult => {
          const selectedTabId =
            this.dataCache.Tabs()[this._selectTab.index]!.Id;

          this.noodleApi
            .noodleApplicationTabsPost(
              NoodleApplicationTabsPostActionEnum.UpdateTab,
              {
                ApplicationId: appResult.Id!,
                TabId: selectedTabId,
              }
            )
            .then(() => {
              this._dialog.close();
            })
            .catch(reason => {
              this.showError(reason);
            });
        })
        .catch(reason => {
          this.showError(reason);
        });

      return;
    }

    this._textFieldApplicationName.reportValidity();
  }

  private showError(reason: ResponseError) {
    reason.response.json().then((value: any) => {
      this.errorText = value.message;
    });
  }

  private dialogClosed(event: Event) {
    if (event.target === event.currentTarget) {
      this.errorText = '';
      this.dispatchEvent(new Event('edit-user-application-dialog-closed'));
    }
  }

  private tabsListTemplate() {
    return html`
      ${this.dataCache
        .Tabs()
        .map(tab => html`<mwc-list-item><b>${tab.Label}</b> </mwc-list-item>`)}
    `;
  }

  render() {
    return html`
      <mwc-dialog
        id="dialog"
        heading="Edit User Application"
        scrimClickAction=""
        @closed=${this.dialogClosed}
      >
        <div class="vertical">
          <mwc-select outlined id="select-tab" label="Tab">
            ${this.tabsListTemplate()}
          </mwc-select>
          <mwc-textfield
            outlined
            id="text-field-application-name"
            minlength="3"
            maxlength="50"
            label="Application name"
            required
            value="${this.application.Name!}"
          >
          </mwc-textfield>
          <mwc-textfield
            outlined
            id="text-field-application-url"
            minlength="6"
            maxlength="246"
            label="Url"
            required
            value="${this.application.Website!}"
          >
          </mwc-textfield>
          <mwc-textfield
            outlined
            id="text-field-application-background"
            minlength="6"
            maxlength="246"
            label="Background Colour"
            required
            value="${this.application.TileBackground!}"
          >
          </mwc-textfield>
          <mwc-textfield
            outlined
            id="text-field-application-icon"
            minlength="6"
            maxlength="246"
            label="Icon"
            required
            value="${this.application.Icon!}"
          >
          </mwc-textfield>
          <img
            width="128px"
            height="128px"
            src="/out-tsc/icons/${this.application.Icon}"
            alt="icon"
          />
          <div>${this.errorText}</div>
        </div>

        <mwc-button
          id="primary-action-button"
          slot="primaryAction"
          @click=${this.primaryButtonClick}
        >
          Confirm
        </mwc-button>
        <mwc-button slot="secondaryAction" dialogAction="close">
          Cancel
        </mwc-button>
      </mwc-dialog>
    `;
  }
}
