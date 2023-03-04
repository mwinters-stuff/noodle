import { html, css, LitElement } from 'lit';
import { query, customElement, state } from 'lit/decorators.js';
import { consume, ContextConsumer } from '@lit-labs/context';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-icon-button';
import '@material/mwc-list';
import '@material/mwc-snackbar';

import * as mwcSnackBar from '@material/mwc-snackbar';

import {
  NoodleApiApi,
  Tab,
  UserSession,
} from '../api/index.js';
import { noodleApiContext, userSessionContext } from '../noodle-context.js';

@customElement('noodle-settings-tabs')
export class NoodleSettingsTabs extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @consume({ context: userSessionContext, subscribe: true })
  @state()
  userSession!: UserSession;

  private _userSession = new ContextConsumer(
    this,
    userSessionContext,
    () => this.refreshTabs(),
    true
  );

  @state()
  tabs: Tab[] = [];

  @state()
  selectedTab: Tab | undefined;

  @state()
  errorMessage = '';

  @query('#error-snack')
  _errorSnack!: mwcSnackBar.Snackbar;

  static styles = css`
    :host {
      display: block;
      border-width: 0;
      width: 100%;
      height: 100%;
    }
    #content {
      margin-top: 8px;
      margin-left: 8px;
      margin-right: 8px;
      margin-bottom: 8px;
    }
    div.ed-buttons {
      display: flex;
      flex-direction: row;
      justify-content: flex-end;
      align-items: center;
    }
    mwc-icon-button[disabled] {
      opacity: 0.5; /* reduce opacity to create a disabled effect */
    }
  `;

  firstUpdated() {}

  refreshTabs() {
    if (this.userSession != null && this.userSession.UserId != null) {
      this.noodleApi
        .noodleTabsGet()
        .then(value => {
          this.tabs = value;
        })
        .catch(reason => {
          this.showError(reason);
        });
    }
  }

  showError(error: string) {
    this.errorMessage = error;
    this._errorSnack.show();
  }


  TabsListTemplate() {
    return html`
      <mwc-list>
            ${this.tabs.map(
              (tab, index, array) =>
                html`<mwc-list-item hasMeta>
                  <span>${tab.Label}</span>
                  <div class="ed-buttons" slot="meta">
                    <mwc-icon-button
                      icon="edit"
                      @click=${() =>
                        this.editTabDialog(
                          tab,
                        )}
                    ></mwc-icon-button>
                    <mwc-icon-button
                      icon="delete"
                      @click=${() => this.deleteTabDialog(tab)}
                    ></mwc-icon-button>
                    <mwc-icon-button
                      icon="arrow_upward"
                      ?disabled=${index === 0}
                      @click=${() =>
                        this.moveTab(tab, index, -1)}
                    ></mwc-icon-button>
                    <mwc-icon-button
                      icon="arrow_downward"
                      ?disabled=${index === array.length - 1}
                      @click=${() =>
                        this.moveTab(tab, index, 1)}
                    ></mwc-icon-button>
                  </div>
                </mwc-list-item> `
            )}
      </mwc-list>
    `;
  }

  private editTabDialog(tab: Tab) {
    // this._editUserApplication.show(application, tabId);
  }

  private deleteTabDialog(tab:Tab) {
    // this._deleteUserApplication.show(userApplication);
  }

  private static swapElements(
    array: Tab[],
    index1: number,
    index2: number
  ): Tab[] {
    const temp = array[index1];
    const a2 = array;
    // Swap the values of the properties of the two items
    a2[index1] = array[index2];
    a2[index2] = temp;
    return a2;
  }

  private moveTab(
    tab: Tab,
    index: number,
    by: number
  ) {
    // const apps = NoodleSettingsTabs.swapElements(
    //   this.getAppsForTab(tabid),
    //   index,
    //   index + by
    // );
    // // update tab indexes.
    // apps.forEach((value: UserApplications, indexInList: number) => {
    //   const apptabid = this.getAppTabIDForAppInTab(tabid, value.ApplicationId!);
    //   if (apptabid > -1) {
    //     this.noodleApi
    //       .noodleApplicationTabsPost(
    //         NoodleApplicationTabsPostActionEnum.UpdateDisplayOrder,
    //         {
    //           Id: apptabid,
    //           DisplayOrder: indexInList,
    //         }
    //       )
    //       .catch(reason => {
    //         this.showError(reason);
    //       });
    //   }
    // });
    this.refreshTabs();
  }

  private showAddTabDialog() {
    // this._addUserApplication.show();
  }

  render() {
    return html`
      <div id="Tabcontent">${this.TabsListTemplate()}</div>
      <mwc-snackbar id="error-snack" labelText="${this.errorMessage}">
        <mwc-icon-button icon="close" slot="dismiss"></mwc-icon-button>
      </mwc-snackbar>
    `;
  }
}
