import { html, css, LitElement } from 'lit';
import { customElement, state } from 'lit/decorators.js';
import { consume } from '@lit-labs/context';

import '@material/mwc-top-app-bar-fixed';
import '@material/mwc-icon-button';
import '@material/mwc-list';

import '@material/mwc-tab-indicator';
import '@shoelace-style/shoelace/dist/components/tab/tab.js';
import '@shoelace-style/shoelace/dist/components/tab-group/tab-group.js';
import '@shoelace-style/shoelace/dist/components/tab-panel/tab-panel.js';

// import { SlTab, SlTabGroup, SlTabPanel } from '@shoelace-style/shoelace/dist/components/tab;
import './noodle-settings-tabs.js'

import {
  NoodleApiApi,
  UserSession,
} from '../api/index.js';
import { DataCache, dataCacheContext, noodleApiContext, userSessionContext } from '../noodle-context.js';


@customElement('noodle-settings')
export class NoodleSettings extends LitElement {
  @consume({ context: noodleApiContext })
  @state()
  noodleApi!: NoodleApiApi;

  @consume({ context: userSessionContext, subscribe: true })
  @state()
  userSession!: UserSession;

  @consume({ context: dataCacheContext })
  @state()
  dataCache!: DataCache;

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
  }



  render() {
    return html`
      <mwc-top-app-bar-fixed id="top-app-bar">
        <mwc-icon-button slot="navigationIcon" icon="arrow_back" @click=${() => window.history.back()}
          ></mwc-icon-button>
        <div slot="title" id="title">Noodle - Settings</div>
      
        <div id="content">
          <sl-tab-group id="tabbar">
            <sl-tab slot="nav" panel="setting-tabs" icon="tabs">Tabs</sl-tab>
            <sl-tab slot="nav" panel="setting-group-apps" icon="apps">Group Apps</sl-tab>
            <sl-tab slot="nav" panel="setting-other" icon="settings">Other Settings</sl-tab>
            <sl-tab-panel name="setting-tabs">
              <noodle-settings-tabs> </noodle-settings-tabs>
            </sl-tab-panel>
            <sl-tab-panel name="setting-group-apps">group apps</sl-tab-panel>
            <sl-tab-panel name="setting-other">other settings</sl-tab-panel>
          </sl-tab-group>
        </div>
      </mwc-top-app-bar-fixed>
    `;
  }
}
