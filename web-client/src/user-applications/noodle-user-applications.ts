import { html, css, LitElement } from 'lit';
import { query, customElement, state } from 'lit/decorators.js';
// eslint-disable-next-line import/no-extraneous-dependencies
import { consume, ContextConsumer } from '@lit-labs/context';

import { setBasePath } from '@shoelace-style/shoelace/dist/utilities/base-path.js';

import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/dialog/dialog.js';


import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-icon-button';
import '@material/mwc-list';
import '@material/mwc-snackbar';

import './noodle-add-user-application.js';
import './noodle-edit-user-application.js';
import './noodle-delete-user-application.js';

import * as mwcSnackBar from '@material/mwc-snackbar';
import * as aua from './noodle-add-user-application.js';
import * as eua from './noodle-edit-user-application.js';
import * as dua from './noodle-delete-user-application.js';

import {
  Application,
  ApplicationTab,
  NoodleApiApi,
  NoodleApplicationTabsPostActionEnum,
  Tab,
  UserApplications,
  UserSession,
} from '../api/index.js';
import { noodleApiContext, userSessionContext } from '../noodle-context.js';

@customElement('noodle-user-applications')
export class NoodleUserApplications extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @consume({ context: userSessionContext, subscribe: true })
  @state()
  userSession!: UserSession;

  private _userSession = new ContextConsumer(
    this,
    userSessionContext,
    () => this.refreshUserApplications(),
    true
  );

  @state()
  tabs: Tab[] = [];

  @state()
  selectedApplication: UserApplications | undefined;

  @state()
  applicationTabs: ApplicationTab[][] = [];

  @state()
  errorMessage = '';

  @state()
  userApplications: UserApplications[] = [];

  @query('#error-snack')
  _errorSnack!: mwcSnackBar.Snackbar;

  @query('#add-user-application')
  _addUserApplication!: aua.NoodleAddUserApplication;

  @query('#edit-user-application')
  _editUserApplication!: eua.NoodleEditUserApplication;

  @query('#delete-user-application')
  _deleteUserApplication!: dua.NoodleDeleteUserApplication;

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

  firstUpdated() {
    // Set the base path to the folder you copied Shoelace's assets to
    setBasePath('node_modules/@shoelace-style/shoelace/dist');
  }

  refreshUserApplications() {
    if (this.userSession != null && this.userSession.UserId != null) {
      this.noodleApi
        .noodleUserApplicationsGet(this.userSession.UserId!)
        .then(value => {
          this.userApplications = value;

          this.noodleApi.noodleTabsGet().then(tabsvalue => {
            this.tabs = tabsvalue;
            this.tabs.forEach(tab => {
              this.noodleApi
                .noodleApplicationTabsGet(tab.Id!)
                .then(appTabData => {
                  this.applicationTabs[tab.Id!] = appTabData;
                  this.requestUpdate();
                })
                .catch(reason => {
                  this.showError(reason);
                });
            });
          });
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

  private getAppsForTab(tabid: number): UserApplications[] {
    const result: UserApplications[] = [];
    if (this.applicationTabs[tabid]) {
      this.applicationTabs[tabid].forEach(appInTab => {
        this.userApplications.forEach(userApp => {
          if (userApp.ApplicationId === appInTab.ApplicationId) {
            result.push(userApp);
          }
        });
      });
    }
    return result;
  }

  private getAppTabIDForAppInTab(tabid: number, applicationId: number): number {
    if (this.applicationTabs[tabid]) {
      const apptab = this.applicationTabs[tabid].find(
        value => value.ApplicationId === applicationId
      );
      if (apptab != null) {
        return apptab.Id!;
      }
    }
    return -1;
  }

  applicationsListTemplate() {
    return html`
      <mwc-list>
        ${this.tabs.map(
          tab => html`<mwc-list-item noninteractive
              ><b>${tab.Label}</b>
            </mwc-list-item>
            <li divider padded role="separator"></li>
            ${this.getAppsForTab(tab.Id!).map(
              (app, index, array) =>
                html`<mwc-list-item graphic="avatar" twoLine hasMeta>
                  <span>${app.Application?.Name}</span>
                  <span slot="secondary">${app.Application?.Website}</span>
                  <img
                    slot="graphic"
                    src="/out-tsc/icons/${app.Application?.Icon}"
                    alt="${app.Application?.Icon}"
                  />
                  <div class="ed-buttons" slot="meta">
                    <mwc-icon-button
                      icon="edit"
                      @click=${() =>
                        this.editUserApplicationDialog(
                          app.Application!,
                          tab.Id!
                        )}
                    ></mwc-icon-button>
                    <mwc-icon-button
                      icon="delete"
                      @click=${() => this.deleteUserApplicationDialog(app)}
                    ></mwc-icon-button>
                    <mwc-icon-button
                      icon="arrow_upward"
                      ?disabled=${index === 0}
                      @click=${() =>
                        this.moveApp(app.Application!, tab.Id!, index, -1)}
                    ></mwc-icon-button>
                    <mwc-icon-button
                      icon="arrow_downward"
                      ?disabled=${index === array.length - 1}
                      @click=${() =>
                        this.moveApp(app.Application!, tab.Id!, index, 1)}
                    ></mwc-icon-button>
                  </div>
                </mwc-list-item> `
            )}`
        )}
      </mwc-list>
    `;
  }

  private editUserApplicationDialog(application: Application, tabId: number) {
    this._editUserApplication.show(application, tabId);
  }

  private deleteUserApplicationDialog(userApplication: UserApplications) {
    this._deleteUserApplication.show(userApplication);
  }

  private static swapElements(
    array: UserApplications[],
    index1: number,
    index2: number
  ): UserApplications[] {
    const temp = array[index1];
    const a2 = array;
    // Swap the values of the properties of the two items
    a2[index1] = array[index2];
    a2[index2] = temp;
    return a2;
  }

  private moveApp(
    application: Application,
    tabid: number,
    index: number,
    by: number
  ) {
    const apps = NoodleUserApplications.swapElements(
      this.getAppsForTab(tabid),
      index,
      index + by
    );
    // update tab indexes.
    apps.forEach((value: UserApplications, indexInList: number) => {
      const apptabid = this.getAppTabIDForAppInTab(tabid, value.ApplicationId!);
      if (apptabid > -1) {
        this.noodleApi
          .noodleApplicationTabsPost(
            NoodleApplicationTabsPostActionEnum.UpdateDisplayOrder,
            {
              Id: apptabid,
              DisplayOrder: indexInList,
            }
          )
          .catch(reason => {
            this.showError(reason);
          });
      }
    });
    this.refreshUserApplications();
  }

  private showAddUserApplicationDialog() {
    this._addUserApplication.show();
  }

  render() {
    return html`
      <mwc-top-app-bar-fixed id="top-app-bar">
        <mwc-icon-button
          slot="navigationIcon"
          icon="arrow_back"
          @click=${() => window.history.back()}
        ></mwc-icon-button>
        <div slot="title" id="title">Noodle - User Applications</div>
        <mwc-icon-button
          icon="add"
          slot="actionItems"
          class="notCenter"
          @click=${() => this.showAddUserApplicationDialog()}
        ></mwc-icon-button>
        <div id="content">${this.applicationsListTemplate()}</div>
      </mwc-top-app-bar-fixed>
      <noodle-add-user-application
        id="add-user-application"
        @add-user-application-dialog-closed=${this.refreshUserApplications}
      ></noodle-add-user-application>

      <noodle-edit-user-application
        id="edit-user-application"
        @edit-user-application-dialog-closed=${this.refreshUserApplications}
      ></noodle-edit-user-application>

      <noodle-delete-user-application
        id="delete-user-application"
        @delete-user-application-dialog-closed=${this.refreshUserApplications}
      ></noodle-delete-user-application>

      <mwc-snackbar id="error-snack" labelText="${this.errorMessage}">
        <mwc-icon-button icon="close" slot="dismiss"></mwc-icon-button>
      </mwc-snackbar>
    `;
  }
}
