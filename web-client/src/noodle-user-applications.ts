import { html, css, LitElement } from 'lit';
import { query, customElement, property, state } from 'lit/decorators.js';
// eslint-disable-next-line import/no-extraneous-dependencies
import { consume, ContextConsumer } from '@lit-labs/context';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-icon-button';
import '@material/mwc-list';
import '@material/mwc-snackbar';
import './noodle-add-user-application.js';
import { Router } from '@vaadin/router';

import * as mwcSnackBar from '@material/mwc-snackbar';
import * as aua from './noodle-add-user-application.js';

import {
  ApplicationTab,
  NoodleApiApi,
  Tab,
  UserApplications,
  UserSession,
} from './api/index.js';
import { noodleApiContext, userSessionContext } from './noodle-context.js';

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

  @property({ attribute: false })
  tabs: Tab[] = [];

  @property({ attribute: false })
  selectedApplication: UserApplications | undefined;

  @property({ attribute: false })
  applicationTabs: ApplicationTab[][] = [];

  @property({ attribute: false })
  errorMessage = '';

  @property({ attribute: false })
  userApplications: UserApplications[] = [];

  @query('#error-snack')
  _errorSnack!: mwcSnackBar.Snackbar;

  @query('#add-user-application')
  _addUserApplication!: aua.NoodleAddUserApplication;

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
  `;

  firstUpdated() {}

  refreshUserApplications() {
    if (this.userSession != null && this.userSession.userId != null) {
      this.noodleApi
        .noodleUserApplicationsGet({ userId: this.userSession.userId! })
        .then(value => {
          this.userApplications = value;

          this.noodleApi.noodleTabsGet().then(tabs => {
            this.tabs = tabs;
            tabs.forEach(tab => {
              this.noodleApi
                .noodleApplicationTabsGet({ tabId: tab.id! })
                .then(appTab => {
                  this.applicationTabs[tab.id!] = appTab;
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
          if (userApp.applicationId === appInTab.applicationId) {
            result.push(userApp);
          }
        });
      });
    }
    return result;
  }

  applicationsListTemplate() {
    return html`
      <mwc-list>
        ${this.tabs.map(
          tab => html`<mwc-list-item><b>${tab.label}</b> </mwc-list-item>
            <li divider padded role="separator"></li>
            ${this.getAppsForTab(tab.id!).map(
              app =>
                html`<mwc-list-item>${app.application?.name}</mwc-list-item> `
            )}`
        )}
      </mwc-list>
    `;
  }

  showAddUserApplicationDialog() {
    this._addUserApplication.show();
  }

  render() {
    return html`
      <mwc-top-app-bar-fixed id="top-app-bar">
        <mwc-icon-button
          slot="navigationIcon"
          icon="arrow_back"
          @click=${() => Router.go('/dash')}
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
        @closed=${this.refreshUserApplications}
      ></noodle-add-user-application>
      <mwc-snackbar id="error-snack" labelText="${this.errorMessage}">
        <mwc-icon-button icon="close" slot="dismiss"></mwc-icon-button>
      </mwc-snackbar>
    `;
  }
}
