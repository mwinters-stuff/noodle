import { html, LitElement } from 'lit';

import '@material/mwc-button';
import { customElement, property, state } from 'lit/decorators.js';
import { consume, ContextConsumer } from '@lit-labs/context';
import { Application } from './api/index.js';
import { DataCache, dataCacheContext } from './noodle-context.js';

@customElement('noodle-app-card')
export class NoodleAppCard extends LitElement {
  @consume({ context: dataCacheContext, subscribe: true })
  @state()
  dataCache!: DataCache;

  private _dataCache = new ContextConsumer(
    this,
    dataCacheContext,
    () => {
      this.application = this.dataCache.GetApplication(this.appId);
      // console.log("App Card: ", this.appId, JSON.stringify(this.application, null, 2))
    },
    true
  );

  @property({ type: Number })
  public appId!: number;

  @state()
  application!: Application;

  render() {
    return html`
      <div>
        <mwc-button outlined label="${this.application?.Name}"> </mwc-button>
        <img
          src="/out-tsc/icons/${this.application?.Icon}"
          alt="${this.application?.Icon}"
          width="64px"
          height="64px"
        />
      </div>
    `;
  }
}
