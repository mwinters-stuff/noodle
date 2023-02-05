import { html, css } from 'lit';
import { query, customElement, property } from 'lit/decorators.js';

// import * as mwcTopAppBarFixed from '@material/mwc-top-app-bar-fixed';
// eslint-disable-next-line import/no-extraneous-dependencies
import '@material/mwc-top-app-bar-fixed';

// import * as mwcDrawer from '@material/mwc-drawer';
import '@material/mwc-drawer';
import '@material/mwc-icon-button';
import '@material/mwc-list';
import '@material/mwc-snackbar';
// eslint-disable-next-line import/no-extraneous-dependencies
import { Router } from '@vaadin/router';

import * as mwcSnackBar from '@material/mwc-snackbar';

import { NoodleAuthenticatedBase } from './noodle-authenticated-base.js';
import { Tab } from './noodleApi.js';

@customElement('noodle-dash')
export class NoodleDash extends NoodleAuthenticatedBase {
  @property({ type: Boolean }) openDrawer = false;

  @property() tabs: Tab[] = [];

  @property() selectedTab: Tab | undefined;

  @property() errorMessage = '';

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

  firstUpdated() {
    super.firstUpdated();
    NoodleAuthenticatedBase.getAuthenticatedApi()
      .noodle.tabsList()
      .then(response => {
        if (response.error != null) {
          this.showError(response.error.message!.toString());
        } else {
          this.tabs = response.data;
          if (this.selectedTab == null && this.tabs.length > 0) {
            this.selectedTab = this.tabs.at(0);
          }
          NoodleAuthenticatedBase.getAuthenticatedApi()
            .NoodleAuthenticatedBase.getAuthenticatedApi()
            .noodle.userAllowedApplicationsList();
        }
      });
  }

  showError(error: string) {
    this.errorMessage = error;
    this._errorSnack.show();
  }

  render() {
    return html`
      <mwc-drawer hasHeader type="modal" ?open=${this.openDrawer}>
        <span slot="title">Noodle</span>
        <span slot="subtitle">App Categories</span>
        <mwc-list activatable>
          ${this.tabs.map(
            tab => html`<mwc-list-item
              ?selected=${this.selectedTab === tab}
              ?activated=${this.selectedTab === tab}
              @click=${() => {
                this.selectedTab = tab;
                this.openDrawer = false;
              }}
              >${tab.Label}
            </mwc-list-item>`
          )}
        </mwc-list>

        <section slot="appContent">
          <mwc-top-app-bar-fixed id="top-app-bar">
            <mwc-icon-button
              slot="navigationIcon"
              icon="menu"
              @click=${() => this.toggleHamburger()}
            ></mwc-icon-button>
            <div slot="title" id="title">
              Noodle - ${this.selectedTab?.Label}
            </div>

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

            <div id="content">
              <mwc-button
                id="prominentButton"
                outlined
                label="test ${this.selectedTab?.Label}"
                @click=${() => this.showError('failed')}
              >
              </mwc-button>
            </div>
          </mwc-top-app-bar-fixed>
        </section>
      </mwc-drawer>
      <mwc-snackbar id="error-snack" labelText="${this.errorMessage}">
        <mwc-icon-button icon="close" slot="dismiss"></mwc-icon-button>
      </mwc-snackbar>
    `;
  }
}
