import { html, css, LitElement } from 'lit';
import { query, customElement, state } from 'lit/decorators.js';
// eslint-disable-next-line import/no-extraneous-dependencies
import * as mwcButton from '@material/mwc-button';
import * as mwcDialog from '@material/mwc-dialog';
import * as mwcTextField from '@material/mwc-textfield';
import * as mwcSelect from '@material/mwc-select';

import { consume, ContextConsumer } from '@lit-labs/context';
import {
  DataCache,
  dataCacheContext,
  noodleApiContext,
  userSessionContext,
} from './noodle-context.js';
import {
  ApplicationTemplate,
  NoodleApiApi,
  NoodleApplicationTabsPostActionEnum,
  UserSession,
} from './api/index.js';

import '@material/mwc-dialog';
import '@material/mwc-button';
import '@material/mwc-textfield';
import '@material/mwc-formfield';
import '@material/mwc-select';

@customElement('noodle-add-user-application')
export class NoodleAddUserApplication extends LitElement {
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

  @query('#select-template')
  _selectTemplate!: mwcSelect.Select;

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

  private _userSession = new ContextConsumer(
    this,
    userSessionContext,
    () => this.Reload(),
    true
  );

  private Reload() {
    // eslint-disable-next-line no-console

    this.noodleApi.noodleAppTemplatesGet({ search: 'A' }).then(value => {
      this._appTemplates = value;
    });
  }

  public show() {
    // this._textFieldApplicationName.value = '';
    this._dialog.show();
  }

  private primaryButtonClick() {
    const isValid = this._textFieldApplicationName.checkValidity();
    if (isValid) {
      let templateAppId: string | undefined;
      if (this._selectTemplate.index > 0) {
        templateAppId =
          this._appTemplates[this._selectTemplate.index - 1].appid;
      }

      this.noodleApi
        .noodleApplicationsPost({
          application: {
            templateAppid: templateAppId,
            name: this._textFieldApplicationName.value,
            website: this._textFieldApplicationUrl.value,
            license: '',
            description: this._textFieldApplicationName.value,
            tileBackground: this._textFieldBackground.value,
            icon: this._textFieldIcon.value,
            enhanced: false,
          },
        })
        .then(appResult => {
          this.noodleApi
            .noodleUserApplicationsPost({
              userApplication: {
                applicationId: appResult.id,
                userId: this._userSession.value?.userId,
              },
            })
            .then(() => {
              const selectedTabId =
                this.dataCache.Tabs()[this._selectTab.index]!.id;

              this.noodleApi
                .noodleApplicationTabsPost({
                  action: NoodleApplicationTabsPostActionEnum.Insert,
                  applicationTab: {
                    applicationId: appResult.id!,
                    tabId: selectedTabId,
                  },
                })
                .then(() => {
                  this._dialog.close();
                });
            });
        });

      return;
    }

    this._textFieldApplicationName.reportValidity();
  }

  private dialogClosed() {
    this.dispatchEvent(new Event('closed'));
  }

  private applicationTemplateSelected() {
    if (this._selectTemplate.index > 0) {
      const template = this._appTemplates[this._selectTemplate.index - 1];
      this._textFieldApplicationName.value = template.name!;
      this._textFieldApplicationUrl.value = template.website!;
      this._textFieldBackground.value = template.tileBackground!;
      this._textFieldIcon.value = template.icon!;
      this._icon = template.icon!;
    }
  }

  private appTemplatesListTemplate() {
    if (this._appTemplates != null) {
      return html`
        <mwc-list-item></mwc-list-item>
        ${this._appTemplates.map(
          at => html`<mwc-list-item><b>${at.name}</b> </mwc-list-item>`
        )}
      `;
    }
    return html``;
  }

  private tabsListTemplate() {
    return html`
      ${this.dataCache
        .Tabs()
        .map(tab => html`<mwc-list-item><b>${tab.label}</b> </mwc-list-item>`)}
    `;
  }

  render() {
    return html`
      <mwc-dialog
        id="dialog"
        heading="Add User Application"
        scrimClickAction=""
        @closed=${this.dialogClosed}
      >
        <div class="vertical">
          <mwc-select
            outlined
            id="select-template"
            @selected=${this.applicationTemplateSelected}
            label="Application Template"
          >
            ${this.appTemplatesListTemplate()}
          </mwc-select>
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
          >
          </mwc-textfield>
          <mwc-textfield
            outlined
            id="text-field-application-url"
            minlength="6"
            maxlength="246"
            label="Url"
            required
          >
          </mwc-textfield>
          <mwc-textfield
            outlined
            id="text-field-application-background"
            minlength="6"
            maxlength="246"
            label="Background Colour"
            required
          >
          </mwc-textfield>
          <mwc-textfield
            outlined
            id="text-field-application-icon"
            minlength="6"
            maxlength="246"
            label="Icon"
            required
          >
          </mwc-textfield>
          <img
            width="128px"
            height="128px"
            src="/out-tsc/icons/${this._icon}"
            alt="icon"
          />
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