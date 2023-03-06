import { customElement, property, state } from 'lit/decorators.js';
import { consume, ContextConsumer } from '@lit-labs/context';
import { Application } from './api/index.js';
import { DataCache, dataCacheContext } from './noodle-context.js';
import { NoodleAppCard } from './common/noodle-app-card.js';
import { Functions } from './common/functions.js';


@customElement('noodle-dash-app-card')
export class NoodleDashAppCard extends NoodleAppCard {

  @consume({ context: dataCacheContext, subscribe: true })
  @state()
  dataCache!: DataCache;

  private _dataCache = new ContextConsumer(
    this,
    dataCacheContext,
    () => {
      if (this.appId) {
        this.application = this.dataCache.GetApplication(this.appId);
        this.updateValues();
      }
    },
    true
  );

  @property({ type: Number })
  public appId!: number;

  @state()
  public application!: Application;

  public updateValues() {
    this.appTitle = this.application.Name!;
    this.appUrl = this.application.Website!;
    this.textColor = Functions.modifyColor(this.application.TextColor!);
    this.backgroundColor = Functions.modifyColor(this.application.TileBackground!);
    this.appIconUrl = this.application.Icon!;
    // super.updateValues();
  }
}

