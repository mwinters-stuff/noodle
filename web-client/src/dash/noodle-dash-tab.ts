import { html, css, LitElement } from 'lit';
import { customElement, state, property } from 'lit/decorators.js';
import { consume, ContextConsumer } from '@lit-labs/context';

import {
  UsersApplicationItem,
} from '../api/index.js';
import { DataCache, dataCacheContext } from '../noodle-context.js';

import './noodle-dash-app-card.js';

@customElement('noodle-dash-tab')
export class NoodleDashTab extends LitElement {

  @property({ type: Number })
  tabId: number = -1;

  @consume({ context: dataCacheContext })
  @state()
  dataCache!: DataCache;

  private _dataCacheConsumer = new ContextConsumer(
    this,
    dataCacheContext,
    () => this.Refresh(),
    true
  );

  @state()
  private _userApplications: UsersApplicationItem[] = [];


  private Refresh() {
    this._userApplications = this.dataCache.GetUserApplicationsForTab(this.tabId);
  }


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
  `

  render() {
    return html`
      ${this._userApplications.map(
        app =>
          html`<noodle-dash-app-card
            appId="${app.Application?.Id}"
          ></noodle-dash-app-card>`
        )}
      `
  }
}