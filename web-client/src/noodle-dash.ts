import { html, css, LitElement } from 'lit';
import { query, customElement, state, property } from 'lit/decorators.js';
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

import './noodle-app-card.js';

import {
  NoodleApiApi,
  Tab,
  UsersApplicationItem,
  UserSession,
} from './api/index.js';
import {
  DataCache,
  dataCacheContext,
  noodleApiContext,
  userSessionContext,
} from './noodle-context.js';

@customElement('noodle-dash')
export class NoodleDash extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @consume({ context: userSessionContext, subscribe: true })
  @state()
  userSession!: UserSession;

  @consume({ context: dataCacheContext })
  @state()
  dataCache!: DataCache;

  @property({ type: Number })
  public tabId?: number;

  private _userSessionConsumer = new ContextConsumer(
    this,
    userSessionContext,
    () => this.RefreshUserApplications(),
    true
  );

  @state()
  openDrawer = false;

  @state()
  selectedTab: Tab | undefined;

  @state()
  errorMessage = '';

  @query('#error-snack')
  _errorSnack!: mwcSnackBar.Snackbar;

  @state()
  private _userApplications: UsersApplicationItem[] = [];

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

  @state()
  private _tabs: Tab[] = [];

  firstUpdated() {
    if (this.dataCache.router && this.dataCache.router!.location.params.tabId) {
      this.tabId = parseInt(
        this.dataCache.router!.location.params.tabId[0],
        10
      );
    }
  }

  toggleHamburger() {
    this.openDrawer = !this.openDrawer;
  }

  private RefreshUserApplications() {
    if (this.userSession != null && this.userSession.UserId != null) {
      this.noodleApi
        .noodleUserAllowedApplicationsGet(this.userSession.UserId!)
        .then(value => {
          this._userApplications = value;
          this.dataCache.SetUserApplications(value);
          // console.log(JSON.stringify(this._userApplications,null,2  ))
        })
        .catch(reason => {
          this.showError(reason);
        });
      this.RefreshTabs();
    }
  }

  private RefreshTabs() {
    // console.log("RefreshTabs")
    this.noodleApi.noodleTabsGet().then(value => {
      this._tabs = value;
      if ((!this.tabId || this.tabId === -1) && this._tabs.length > 0) {
        // console.log("Dash RefreshTabs Redirect")
        Router.go(`/dash/${this._tabs[0].Id}`);
      }

      this._tabs.forEach(tab => {
        if (this.tabId === tab.Id) {
          this.selectedTab = tab;
        }
      });
    });
  }

  showError(error: string) {
    this.errorMessage = error;
    this._errorSnack.show();
  }

  tabListTemplate() {
    return html`
      <mwc-list activatable>
        ${this._tabs.map(
          tab => html`<mwc-list-item
              ?selected=${this.selectedTab === tab}
              ?activated=${this.selectedTab === tab}
              @click=${() => {
                this.selectedTab = tab;
                Router.go(`/dash/${tab.Id}`);
                this.openDrawer = false;
              }}
              >${tab.Label}
            </mwc-list-item>
            <li divider padded role="separator"></li>`
        )}
      </mwc-list>
    `;
  }

  appListTemplate() {
    return html`
      ${this._tabs.map(
        tab =>
          html`
            <div id="tab${tab.Id}" ?hidden=${tab.Id !== this.tabId}>
              ${this._userApplications
                .filter(value => value.TabId === tab.Id)
                .map(
                  app =>
                    html`<noodle-app-card
                      appId="${app.Application?.Id}"
                    ></noodle-app-card>`
                )}
            </div>
          `
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
        <div slot="title" id="title">Noodle - ${this.selectedTab?.Label}</div>

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
