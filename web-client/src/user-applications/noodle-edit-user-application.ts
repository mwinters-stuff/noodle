import { customElement } from 'lit/decorators.js';

import {
  Application,
  NoodleApplicationsPostActionEnum,
  NoodleApplicationTabsPostActionEnum,
} from '../api/index.js';

import { NoodleUserApplicationDialog } from './noodle-user-application-dialog.js';

@customElement('noodle-edit-user-application')
export class NoodleEditUserApplication extends NoodleUserApplicationDialog {

  public show(application: Application, tabId: number) {
    this.application = application;
    this.tabId = tabId.toString();
    this.showDialog('Edit User Application');
  }

  protected primaryButtonClick() {
    this.errorText = '';
    const isValid = this._textFieldApplicationName.checkValidity();
    if (isValid) {
      this.application.Name = this._textFieldApplicationName.value;
      this.application.Website = this._textFieldApplicationUrl.value;
      this.application.Description = this._textFieldApplicationName.value;
      this.application.TileBackground = this._colorPickerBackground.value;
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
          const selectedTabId = this._selectTab.value.toString();

          this.noodleApi
            .noodleApplicationTabsPost(
              NoodleApplicationTabsPostActionEnum.UpdateTab,
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

      return;
    }

    this._textFieldApplicationName.reportValidity();
  }

}
