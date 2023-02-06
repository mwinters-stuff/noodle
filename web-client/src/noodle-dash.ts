import { html, css, LitElement } from 'lit';
import { query, customElement, state } from 'lit/decorators.js';
// eslint-disable-next-line import/no-extraneous-dependencies
import { consume, ContextConsumer } from '@lit-labs/context';

import '@material/mwc-top-app-bar-fixed';

import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-list';
import '@material/mwc-snackbar';
import '@material/mwc-fab';

import { Router } from '@vaadin/router';

import * as mwcSnackBar from '@material/mwc-snackbar';

import {
  NoodleApiApi,
  Tab,
  UsersApplicationItem,
  UserSession,
} from './api/index.js';
import { noodleApiContext, userSessionContext } from './noodle-context.js';

@customElement('noodle-dash')
export class NoodleDash extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @consume({ context: userSessionContext, subscribe: true })
  @state()
  userSession!: UserSession;

  private _userSession = new ContextConsumer(
    this,
    userSessionContext,
    () => this.Refresh(),
    true
  );

  @state()
  openDrawer = false;

  @state()
  tabs: Tab[] = [];

  @state()
  selectedTab: Tab | undefined;

  @state()
  errorMessage = '';

  @state()
  userApplications: UsersApplicationItem[] = [];

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
  `;

  toggleHamburger() {
    this.openDrawer = !this.openDrawer;
  }

  private Refresh() {
    this.noodleApi
      .noodleTabsGet()
      .then(value => {
        this.tabs = value;
        if (this.selectedTab == null && this.tabs.length > 0) {
          this.selectedTab = this.tabs.at(0);
        }

        this.RefreshUserApplications();
      })
      .catch(reason => {
        this.showError(reason);
      });
  }

  private RefreshUserApplications() {
    if (this.userSession != null && this.userSession.userId != null) {
      this.noodleApi
        .noodleUserAllowedApplicationsGet({ userId: this.userSession.userId! })
        .then(value => {
          this.userApplications = value;
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

  tabListTemplate() {
    return html`
      <mwc-list activatable>
        ${this.tabs.map(
          tab => html`<mwc-list-item
              ?selected=${this.selectedTab === tab}
              ?activated=${this.selectedTab === tab}
              @click=${() => {
                this.selectedTab = tab;
                this.openDrawer = false;
              }}
              >${tab.label}
            </mwc-list-item>
            <li divider padded role="separator"></li>`
        )}
      </mwc-list>
    `;
  }

  appListTemplate() {
    return html`
      ${this.userApplications
        .filter(value => value.tabId === this.selectedTab?.id)
        .map(
          app =>
            html`<mwc-button
              outlined
              label="${app.application?.name}"
            ></mwc-button>`
        )}
    `;
  }

  topBarTemplate() {
    return html`
      <mwc-top-app-bar-fixed id="top-app-bar">
        <mwc-icon-button
          slot="navigationIcon"
          icon="menu"
          @click=${() => this.toggleHamburger()}
        ></mwc-icon-button>
        <div slot="title" id="title">Noodle - ${this.selectedTab?.label}</div>

        <mwc-icon-button
          icon="apps"
          slot="actionItems"
          class="notCenter"
          @click=${() => Router.go('/user-applications')}
        ></mwc-icon-button>

        <mwc-icon-button
          icon="settings"
          slot="actionItems"
          class="notCenter"
        ></mwc-icon-button>
        <mwc-icon-button
          icon="logout"
          slot="actionItems"
          class="notCenter"
          @click=${() => Router.go('/logout')}
        >
        </mwc-icon-button>

        <div id="content">${this.appListTemplate()}</div>
      </mwc-top-app-bar-fixed>
    `;
  }

  render() {
    return html`
      <mwc-drawer hasHeader type="modal" ?open=${this.openDrawer}>
        <span slot="title">Noodle</span>
        <span slot="subtitle">App Categories</span>
        ${this.tabListTemplate()}

        <section slot="appContent">${this.topBarTemplate()}</section>
      </mwc-drawer>
      <mwc-snackbar id="error-snack" labelText="${this.errorMessage}">
        <mwc-icon-button icon="close" slot="dismiss"></mwc-icon-button>
      </mwc-snackbar>
    `;
  }
}
