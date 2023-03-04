/* eslint-disable no-console */
import { LitElement, html, css } from 'lit';
import { property, customElement, query } from 'lit/decorators.js';

import { Router } from '@vaadin/router';

import '@shoelace-style/shoelace/dist/themes/light.styles.js';
import '@shoelace-style/shoelace/dist/themes/dark.styles.js';

import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/input/input.js';

import SlInput from '@shoelace-style/shoelace/dist/components/input/input';
import { NoodleAuthApi, UserLogin } from './api/index.js';

@customElement('noodle-login')
export class NoodleLogin extends LitElement {
  @property({ type: String }) header = 'Noodle';

  @query('#username')
  _usernameField!: SlInput;

  @query('#password')
  _passwordField!: SlInput;

  @property({ type: String }) _errorMessage: string = '';

  static styles = css`
    :host {
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: flex-center;
      max-width: 960px;
      margin: 0 auto;
      text-align: center;
      background-color: var(--noodle-web-background-color);
    }

    main {
      flex-grow: 1;
    }
    div.middle {
      max-width: 400px;
      display: flex;
      align-items: center;
      flex-direction: column;
    }
    sl-input{
      text-align: start;
      margin-bottom: 16px;
    }
    // mwc-textfield {
    //   margin-bottom: 16px;
    // }
    div.error {
      margin-top: 10px;
      color: red;
    }
  `;

  keyEvent(ev: KeyboardEvent) {
    if (ev.key === 'Enter') {
      this.login();
    }
  }

  login() {
    this._errorMessage = '';
    const api = new NoodleAuthApi();

    if (!this._usernameField.checkValidity()) {
      this._usernameField.reportValidity();
      this._usernameField.focus();
      return;
    }
    if(!this._passwordField.checkValidity()) {
      this._usernameField.reportValidity();
      this._passwordField.focus();
      return;
    }

    const userLogin: UserLogin = {
      Username: this._usernameField.value,
      Password: this._passwordField.value,
    };

    api
      .authAuthenticatePost(userLogin)
      .then(value => {
        const expires = value.Expires?.toUTCString();
        document.cookie = `noodle-auth=${value.Token}; expires=${expires}; Secure`;
        Router.go('/dash/-1');
      })
      .catch(reason => {
        this._usernameField.focus();
        if (reason.response.statusText) {
          console.error(reason.response.statusText);
          this._errorMessage = reason.response.statusText;
        } else {
          console.error(reason);
          this._errorMessage = reason;

        }
      });
  }



  render() {
    return html`
      <main>
        <div align="center" class="middle">
          <img width="250px" height="250px" src="../../assets/noodle-icon.svg" alt="Noodle Logo"></img>
          <h1>${this.header}</h1>
          <div align="left">
            <sl-input id="username" autofocus minlength="3" maxlength="64" label="Username" required @keyup=${(ev:
              KeyboardEvent)=> this.keyEvent(ev)}>
            </sl-input>
      
            <sl-input id="password" minlength="3" maxlength="64" label="Password" required type="password" password-toggle
              @keyup=${(ev: KeyboardEvent)=> this.keyEvent(ev)}>
            </sl-input>
          </div>
          <sl-button id="login-button" variant="primary" outline @click=${()=>
                  this.login()}>
            Login
          </sl-button>
          <div class="error">${this._errorMessage}</div>
        </div>
      </main>
    `;
  }
}
