import { css, html, LitElement } from 'lit';
import { query, state } from 'lit/decorators.js';

import './noodle-user-application-dialog-app-card.js';

import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/input/input.js';
import '@shoelace-style/shoelace/dist/components/dialog/dialog.js';
import '@shoelace-style/shoelace/dist/components/select/select.js';
import '@shoelace-style/shoelace/dist/components/option/option.js';
import '@shoelace-style/shoelace/dist/components/color-picker/color-picker.js';
import '@shoelace-style/shoelace/dist/components/icon/icon.js';
import '@shoelace-style/shoelace/dist/components/icon-button/icon-button.js';

import SlInput from '@shoelace-style/shoelace/dist/components/input/input';
import SlDialog from '@shoelace-style/shoelace/dist/components/dialog/dialog';
import SlSelect from '@shoelace-style/shoelace/dist/components/select/select';
import SlColorPicker from '@shoelace-style/shoelace/dist/components/color-picker/color-picker';

import { consume, } from '@lit-labs/context';
import {
  DataCache,
  dataCacheContext,
  noodleApiContext,
  userSessionContext,
} from '../noodle-context.js';
import {
  Application,
  ApplicationTemplate,
  NoodleApiApi,
  ResponseError,
  Tab,
  UserSession,
} from '../api/index.js';
import { Functions } from '../common/functions.js';
import { NoodleUserApplicationDialogAppCard } from './noodle-user-application-dialog-app-card.js';

export abstract class NoodleUserApplicationDialog extends LitElement {
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
  application!: Application;

  @state()
  tabId!: string;

  @state()
  _appTemplates!: ApplicationTemplate[];

  @state()
  _icon!: string;

  @state()
  errorText: string = '';

  @state()
  _tabs!: Tab[];

  @state()
  _title: string = "";

  @state()
  _backgroundColor: string = Functions.DarkColor;

  @state()
  _textColor: string = Functions.LightColor;

  @state()
  _appTemplateId: string = "";

  @state()
  _appName: string = "";
  
  @state()
  _appUrl: string = "";

  @state()
  _iconName: string = "";

  @query('#dialog')
  _dialog!: SlDialog;

  @query('#text-field-application-name')
  _textFieldApplicationName!: SlInput;

  @query('#text-field-application-url')
  _textFieldApplicationUrl!: SlInput;

  @query('#text-field-application-background')
  _colorPickerBackground!: SlColorPicker;

  @query('#text-field-text-color')
  _colorPickerText!: SlColorPicker;

  @query('#text-field-application-icon')
  _textFieldIcon!: SlInput;

  @query('#select-template')
  _selectTemplate!: SlSelect;

  @query('#select-tab')
  _selectTab!: SlSelect;

  @query('#icon-file-select')
  _iconFileSelect!: HTMLInputElement;

  @query('#app-card')
  _appCard!: NoodleUserApplicationDialogAppCard;

  private Reload() {
    this._tabs = this.dataCache.Tabs()
    this.noodleApi.noodleAppTemplatesGet('A').then(value => {
      this._appTemplates = value;
      if(this.application){
        this._selectTemplate.value = this.application.TemplateAppid || "";
      }
    });
    
  }

  public showDialog(title: string) {
    this._backgroundColor = Functions.DarkColor;
    this._textColor = Functions.LightColor;
    
    this._title = title;
    this.Reload();
    if(this.application){
      this._appTemplateId = this.application.TemplateAppid || "";
      this._appName = this.application.Name || "";
      this._appUrl = this.application.Website || "";
      this._iconName = this.application.Icon || "";

      this._backgroundColor = Functions.modifyColor(this.application.TileBackground!);
      this._textColor = Functions.modifyColor(this.application.TextColor!);
    }
    this.initAppCard();
    this._dialog.show();
  }

  firstUpdated() {
    this.initAppCard();
  }

  protected abstract primaryButtonClick(): void;

  private applicationTemplateSelected() {
    if (this._selectTemplate.value != "") {
      const template = this._appTemplates.find( value => value.Appid === this._selectTemplate.value);
      if(template){
        this._textFieldApplicationName.value = template.Name!;
        this._textFieldApplicationUrl.value = template.Website!;
        this._backgroundColor = Functions.modifyColor(template.TileBackground!);
        if(template.TileBackground! === "dark"){
          this._textColor = Functions.modifyColor("light");
        }else {
          this._textColor = Functions.modifyColor("dark");
        }
        this._textFieldIcon.value = template.Icon!;
        this._icon = template.Icon!;

        this.updateAppCard();
      }
    }
  }

  private appTemplatesListTemplate() {
    if (this._appTemplates != null) {
      return html`
        ${this._appTemplates.map(
          at => html`<sl-option value="${at.Appid}">${at.Name}</sl-option>`
        )}
      `;
    }
    return html``;
  }

  private tabsListTemplate() {
    if (this._tabs != null) {
      return html`
      ${this._tabs
        .map(tab => html`<sl-option value="${tab.Id}">${tab.Label}</sl-option>`)}
    `;
    }
    return html``;
  }

  private dialogRequestClose(event: CustomEvent){
    if (event.detail.source === 'overlay') {
      event.preventDefault();
    }
  }

  private dialogClosed(event: CustomEvent){
    if(event.target == this._dialog){
      this.errorText = '';
      this.dispatchEvent(new Event('add-user-application-dialog-closed'));
    }
  }

  private secondaryButtonClick(){
    this._dialog.hide();
  }

  private fileUploadClick(){
    this._iconFileSelect.click();

  }

  private initAppCard(){
    this._appCard.appTitle = this._appName;
    this._appCard.appIconUrl = this._iconName;
    this._appCard.appUrl = this._appUrl;
    this._appCard.backgroundColor = this._backgroundColor;
    this._appCard.textColor = this._textColor;
    
  }

  private updateAppCard(){
    this._appCard.appTitle = this._textFieldApplicationName.value;
    this._appCard.appIconUrl = this._textFieldIcon.value;
    this._appCard.appUrl = this._textFieldApplicationUrl.value;
    this._appCard.backgroundColor = this._colorPickerBackground.value;
    this._appCard.textColor = this._colorPickerText.value;
    
  }

  private handleIconFileSelectChange(event: Event){
    if (this._iconFileSelect!.files && this._iconFileSelect!.files!.length == 1) {
      
      const file = this._iconFileSelect!.files![0];
      this._textFieldIcon.value = file.name;
      this._appCard.uploadIcon(file);
      this.noodleApi.noodleUploadIconPost(file).catch(reason => {
        Functions.showWebResponseError(reason);
      });
    }
  }

  static styles = css`
  .no-close::part(close-button) {
    visibility: hidden;
  }
`

  render() {
    return html`
    <input
      type="file"
      id="icon-file-select"
      multiple
      accept="image/*"
      style="display:none"
      @change=${this.handleIconFileSelectChange} />

      <sl-dialog class="no-close"
        id="dialog"
        label="${this._title}"
        @sl-request-close=${this.dialogRequestClose}
        @sl-hide=${this.dialogClosed}
        style="--width: 800px;">
          <sl-select
            id="select-template"
            clearable
            @sl-change=${this.applicationTemplateSelected}
            label="Application Template"
            value="${this._appTemplateId}">
            ${this.appTemplatesListTemplate()}
          </sl-select>
          <sl-select id="select-tab" label="Tab" required value="${this.tabId}">
            ${this.tabsListTemplate()}
          </sl-select>
          <sl-input
            id="text-field-application-name"
            minlength="3"
            maxlength="50"
            label="Application name"
            required
            @sl-input=${this.updateAppCard}
            value="${this._appName}">
          </sl-input>
          <sl-input
            id="text-field-application-url"
            minlength="6"
            maxlength="246"
            label="Url"
            required
            type="url"
            inputmode="url"
            @sl-input=${this.updateAppCard}
            value="${this._appUrl}">
          </sl-input>
          <div style="display: flex; flex-direction: column;">
            <label part="form-control-label" class="form-control__label" for="input" aria-hidden="false">
                Background Colour
            </label>
            <sl-color-picker
              id="text-field-application-background"
              label="Background Colour"
              format="hex"
              swatches="rgb(22,33,37); rgb(250,250,250);"
              @sl-change=${this.updateAppCard}
              value="${this._backgroundColor}">
            </sl-color-picker>
          </div>
          <div style="display: flex; flex-direction: column;">
            <label part="form-control-label" class="form-control__label" for="input" aria-hidden="false">
                Text Colour
            </label>
            <sl-color-picker
              id="text-field-text-color"
              label="Text Colour"
              format="hex"
              swatches="rgb(22,33,37); rgb(250,250,250);"
              @sl-change=${this.updateAppCard}
              value="${this._textColor}">
            </sl-color-picker>
          </div>
          <sl-input
            id="text-field-application-icon"
            minlength="6"
            maxlength="246"
            label="Icon"
            required
            value="${this._iconName}">
            <sl-icon-button name="file-earmark-arrow-up" slot="suffix" @click=${this.fileUploadClick} ></sl-icon-button>
          </sl-input>
          <noodle-user-application-dialog-app-card id="app-card">
          </noodle-user-application-dialog-app-card>
          <div>${this.errorText}</div>
        <sl-button id="primary-action-button" slot="footer" variant="primary" @click=${this.primaryButtonClick}>OK</sl-button>
        <sl-button slot="footer" variant="default" @click=${this.secondaryButtonClick}>Cancel</sl-button>

      </sl-dialog>
    `;
  }
}
