import { customElement } from 'lit/decorators.js';

import {
  NoodleApplicationsPostActionEnum,
  NoodleApplicationTabsPostActionEnum,
} from '../api/index.js';

import { NoodleUserApplicationDialog } from './noodle-user-application-dialog.js';

@customElement('noodle-add-user-application')
export class NoodleAddUserApplication extends NoodleUserApplicationDialog {

  public show() {
    this.showDialog('Add User Application');
  }

  protected primaryButtonClick() {

    this.errorText = '';
    if (!this._selectTab.checkValidity()) {
      this._selectTab.reportValidity();
      this._selectTab.focus();
      return;
    }

    if (!this._textFieldApplicationName.checkValidity()) {
      this._textFieldApplicationName.reportValidity();
      this._textFieldApplicationName.focus();
      return;
    }


    if (!this._textFieldApplicationUrl.checkValidity()) {
      this._textFieldApplicationUrl.reportValidity();
      this._textFieldApplicationUrl.focus();
      return;
    }

    if (!this._colorPickerBackground.checkValidity()) {
      this._colorPickerBackground.reportValidity();
      this._colorPickerBackground.focus();
      return;
    }

    if (!this._textFieldIcon.checkValidity()) {
      this._textFieldIcon.reportValidity();
      this._textFieldIcon.focus();
      return;
    }



    let templateAppId: string | undefined = this._selectTemplate.value.toString();
    if (templateAppId === "") {
      templateAppId = undefined;
    }

    this.noodleApi
      .noodleApplicationsPost(NoodleApplicationsPostActionEnum.Insert, {
        TemplateAppid: templateAppId,
        Name: this._textFieldApplicationName.value,
        Website: this._textFieldApplicationUrl.value,
        License: '',
        Description: this._textFieldApplicationName.value,
        TileBackground: this._colorPickerBackground.value,
        Icon: this._textFieldIcon.value,
        Enhanced: false,
      })
      .then(appResult => {
        this.noodleApi
          .noodleUserApplicationsPost({
            ApplicationId: appResult.Id,
            UserId: this.userSession.UserId,
          })
          .then(() => {
            const selectedTabId = this._selectTab.value.toString();

            this.noodleApi
              .noodleApplicationTabsPost(
                NoodleApplicationTabsPostActionEnum.Insert,
                {
                  ApplicationId: appResult.Id!,
                  TabId: Number.parseInt(selectedTabId, 10),
                }
              )
              .then(() => {
                this._dialog.hide();
              })
              .catch(reason => {
                this.showError(reason);
              });
          })
          .catch(reason => {
            this.showError(reason);
          });
      });

  }
}
