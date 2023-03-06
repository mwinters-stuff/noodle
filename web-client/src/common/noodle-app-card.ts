import { html, css, LitElement, PropertyValues } from 'lit';

import '@material/mwc-button';
import { property, query } from 'lit/decorators.js';

export abstract class NoodleAppCard extends LitElement {

  @query("#item")
  _item!: HTMLDivElement

  @query("#title")
  _title!: HTMLDivElement

  @query("#image")
  public imageElement!: HTMLImageElement

  @property()
  appTitle: string = '';

  @property()
  appIconUrl: string = '';

  @property()
  textColor: string = '';

  @property()
  backgroundColor: string = '';

  @property()
  appUrl: string = '';

  updated(changedProperties: PropertyValues<this>) {
    if (this._item && changedProperties.has("backgroundColor")) {
      this._item.style.backgroundColor = this.backgroundColor;
    }
    if (this._title && changedProperties.has("textColor")) {
      this._title.style.color = this.textColor;
    }

  }

  static styles = css`
    .item-container {
      position: relative;
    }
    .item {
      align-items: center;
      -webkit-background-clip: padding-box;
      background-clip: padding-box;
      background-image: linear-gradient(90deg,hsla(0,0%,100%,0),hsla(0,0%,100%,.25));
      border: 1px solid #4a4a4a;
      border: 1px solid rgba(76,76,76,.4);
      border-radius: 6px;
      color: #fff;
      display: flex;
      flex: 0 0 280px;
      height: 90px;
      margin: 20px;
      outline: 1px solid transparent;
      overflow: hidden;
      padding: 15px 15px 15px 15px;
      position: relative;
      transition: all .35s ease-in-out; 
      width: 280px;
    }

    .app-icon-container {
      align-items: center;
      display: flex;
      flex: 0 0 60px;
      height: 60px;
      justify-content: center;
      margin-right: 15px;
      width: 60px;
    }
    .app-icon {
      display: block;
      max-height: 60px;
      max-width: 60px;
    }
    img {
      border: 0;
    }
    .item .details {
      width: 100%;
    }
    .item .title {
      font-size: 16px;
    }
  `

  render() {
    return html`
      <section class="item-container">
        <div id="item" class="item">
          <div class="app-icon-container">
            <img id="image" class="app-icon" src="/out-tsc/icons/${this.appIconUrl}" alt="${this.appIconUrl}"></img>
          </div>
          <div class="details">
            <div id="title" class="title">${this.appTitle}</div>
          </div>
        </div>
      </section>
    `;
  }
}

