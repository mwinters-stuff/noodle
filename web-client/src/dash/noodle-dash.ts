import { html, css, LitElement } from 'lit';
import { customElement, state, property, query } from 'lit/decorators.js';
import { consume, ContextConsumer } from '@lit-labs/context';

import '@material/mwc-top-app-bar-fixed';

import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-list';

import '@shoelace-style/shoelace/dist/components/tab/tab.js';
import '@shoelace-style/shoelace/dist/components/tab-group/tab-group.js';
import '@shoelace-style/shoelace/dist/components/tab-panel/tab-panel.js';

import SlTabGroup from '@shoelace-style/shoelace/dist/components/tab-group/tab-group';
// import SlTabPanel from '@shoelace-style/shoelace/dist/components/tab-panel/tab-panel';

import { Router } from '@vaadin/router';



import './noodle-dash-tab.js';

import {
  NoodleApiApi,
  Tab,
  UserSession,
} from '../api/index.js';
import {
  DataCache,
  dataCacheContext,
  noodleApiContext,
  userSessionContext,
} from '../noodle-context.js';
import { Functions } from '../common/functions.js';

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

  @state()
  private _selectedTabIndex: number = 0;

  @state()
  private _selectedTabLabel: string = "";

  firstUpdated() {
    if (this.dataCache.router && this.dataCache.router!.location.params.tabId) {
      this.tabId = parseInt(
        this.dataCache.router!.location.params.tabId[0],
        10
      );
    }
  }

  private RefreshUserApplications() {
    if (this.userSession != null && this.userSession.UserId != null) {
      this.noodleApi
        .noodleUserAllowedApplicationsGet(this.userSession.UserId!)
        .then(value => {
          this.dataCache.SetUserApplications(value);
        })
        .catch(reason => {
          Functions.showWebResponseError(reason);
        });
      this.RefreshTabs();
    }
  }

  private RefreshTabs() {
    // console.log("RefreshTabs")
    this.noodleApi.noodleTabsGet().then(value => {
      this._tabs = value;

      this._tabs.forEach((tab: Tab, index: number) => {
        if (this.tabId === tab.Id) {
          this._selectedTabLabel = tab.Label!;
          this._selectedTabIndex = index;
        }
      });

      if ((!this.tabId || this.tabId === -1) && this._tabs.length > 0) {
        Router.go(`/dash/${this._tabs[0].Id}`);
      }

    });
  }

  // @query("#tabgroup")
  // _tabBar!: SlTabGroup

  tabShow(event: CustomEvent){
    const tabId = Number.parseInt(event.detail.name,10);
    const tab = this._tabs.find((value: Tab) => value.Id === tabId)
    this._selectedTabLabel = tab?.Label || "";
    // this._tabBar.show(event.detail.name)
    history.replaceState(null,"",`/dash/${tabId}`)
  }

  tabListTemplate() {
    return html`
    <sl-tab-group id="tabbar" 
     @sl-tab-show=${this.tabShow}
    >
      ${this._tabs.map((tab: Tab) => 
        html`
          <sl-tab slot="nav" panel="${tab.Id}" ?active="${this.tabId === tab.Id}">${tab.Label}</sl-tab>
          <sl-tab-panel name="${tab.Id}" ?active="${this.tabId === tab.Id}">
            <noodle-dash-tab tabId="${tab.Id}"></noodle-dash-tab>
          </sl-tab-panel>
        `
      )}
    </sl-tab-group>
    `;
  }

  topBarTemplate() {
    return html`
      <mwc-top-app-bar-fixed id="top-app-bar">
      
        <div slot="title" id="title">Noodle - ${this._selectedTabLabel}</div>
      
        <mwc-icon-button icon="apps" slot="actionItems" class="notCenter" @click=${()=> Router.go('/user-applications')}
          ></mwc-icon-button>
      
        <mwc-icon-button icon="settings" slot="actionItems" class="notCenter" @click=${()=> Router.go('/settings')}
          ></mwc-icon-button>
        <mwc-icon-button icon="logout" slot="actionItems" class="notCenter" @click=${()=> Router.go('/logout')}
          >
        </mwc-icon-button>
      
        <div id="content">
          ${this.tabListTemplate()}
        </div>
      </mwc-top-app-bar-fixed>
    `;
  }

  render() {
    return html`${this.topBarTemplate()}`;
  }
}
