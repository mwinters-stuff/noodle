/* eslint-disable no-console */
import { LitElement, html, css } from 'lit';
import { property, customElement, query } from 'lit/decorators.js';

// eslint-disable-next-line import/no-extraneous-dependencies
import { Router } from '@vaadin/router';

import '@material/mwc-button';
import '@material/mwc-textfield';
import * as mwcTextfield from '@material/mwc-textfield';
import {
  Api,
  UserLogin,
  UserSession,
  Error,
  HttpResponse,
} from './noodleApi.js';

@customElement('noodle-login')
export class NoodleLogin extends LitElement {
  @property({ type: String }) header = 'Noodle';

  @query('#username')
  _usernameField!: mwcTextfield.TextField;

  @query('#password')
  _passwordField!: mwcTextfield.TextField;

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
    mwc-textfield {
      margin-bottom: 16px;
    }
    div.error {
      margin-top: 10px;
      color: red;
    }
  `;

  // firstUpdated() {

  // }

  login() {
    this._errorMessage = '';

    const api = new Api();
    const userLogin: UserLogin = {
      username: this._usernameField.value,
      password: this._passwordField.value,
    };
    api.auth
      .authenticateCreate(userLogin)
      .then((value: HttpResponse<UserSession, Error>) => {
        if (value.error != null) {
          console.error(value.error);
          this._errorMessage = value.error.message ?? 'Unknown Error';
        } else if (value.ok) {
          const expires = new Date(
            Date.parse(value.data.Expires!)
          ).toUTCString();

          document.cookie = `noodle-auth=${value.data.Token}; expires=${expires}; Secure`;
          Router.go('/dash');
        }
      })
      .catch(reason => {
        console.error(reason.error.message);
        this._errorMessage = reason.error.message;
      });
  }

  render() {
    return html`
      <main>
        <div align="center" class="middle">
          <img width="250px" height="250px" src="../../assets/noodle-icon.svg" alt="Noodle Logo"></img>
          <h1>${this.header}</h1>
          <mwc-textfield outlined id="username" minlength="3" maxlength="64" label="Username" required>
          </mwc-textfield>
      
          <mwc-textfield outlined id="password" minlength="3" maxlength="64" label="Password" required type="password">
          </mwc-textfield>
      
          <mwc-button id="login-button" raised slot="primaryAction" @click=${() =>
            this.login()}>
            Login
          </mwc-button>
          <div class="error">${this._errorMessage}</div>
        </div>
      </main>
    `;
  }
}
