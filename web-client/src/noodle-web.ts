/* eslint-disable no-console */
import { html, css, LitElement } from 'lit';
import { Commands, Context, Router } from '@vaadin/router';
import { customElement, state } from 'lit/decorators.js';

import './noodle-login.js';
import './noodle-dash.js';
import './noodle-user-applications.js';

// eslint-disable-next-line import/no-extraneous-dependencies
import { provide } from '@lit-labs/context';
import {
  AuthSessionGetRequest,
  Configuration,
  ConfigurationParameters,
  NoodleApiApi,
  NoodleAuthApi,
  UserSession,
} from './api/index.js';
import {
  authApiContext,
  noodleApiContext,
  userSessionContext,
} from './noodle-context.js';

@customElement('noodle-web')
export class NoodleWeb extends LitElement {
  @provide({ context: noodleApiContext })
  @state()
  noodleApi: NoodleApiApi;

  @provide({ context: authApiContext })
  @state()
  authApi: NoodleAuthApi;

  @provide({ context: userSessionContext })
  @state()
  userSession: UserSession = {};

  static styles = css`
    :host {
      display: block;
      border-width: 0;
      width: 100%;
      height: 100%;
    }
    main {
      width: 100%;
      height: 100%;
      border-width: 0;
      box-sizing: border-box;
    }
  `;

  activeRoute: string = '';

  params: string = '';

  query: string = '';

  data: string = '';

  constructor() {
    super();
    const parameters: ConfigurationParameters = {
      apiKey: NoodleWeb.apiKey,
    };
    const config = new Configuration(parameters);

    this.noodleApi = new NoodleApiApi(config);
    this.authApi = new NoodleAuthApi(config);

    const params: AuthSessionGetRequest = {
      token: NoodleWeb.getAuthToken(),
    };

    this.authApi
      .authSessionGet(params)
      .then(value => {
        this.userSession = value;
      })
      .catch(reason => {
        console.log('GetUserSession Error: ', reason);
      });
  }

  private static apiKey(name: string): string {
    if (name === 'X-Token') {
      return NoodleWeb.getAuthToken() || '';
    }
    return '';
  }

  private static getAuthToken(): string {
    return (
      document.cookie
        .split('; ')
        .find(row => row.startsWith('noodle-auth='))
        ?.split('=')[1] ?? ''
    );
  }

  public static IsAuthenticated(): boolean {
    return NoodleWeb.checkACookieExists('noodle-auth');
  }

  private static checkACookieExists(cookieName: string): boolean {
    if (
      document.cookie
        .split(';')
        .some(item => item.trim().startsWith(`${cookieName}=`))
    ) {
      return true;
    }
    return false;
  }

  async logout(context: Context, commands: Commands) {
    try {
      await this.authApi.authLogoutGet();
    } catch (e) {
      console.error('Logout Error: ', e);
    }

    document.cookie =
      'noodle-auth=; expires=Thu, 01 Jan 1970 00:00:00 GMT; Secure';

    return commands.redirect('/login'); // pass to the next route in the list
  }

  firstUpdated() {
    const router = new Router(this.shadowRoot!.querySelector('main'));
    router.setRoutes([
      { path: '/', component: 'noodle-dash' },
      { path: '/dash', component: 'noodle-dash' },
      { path: '/login', component: 'noodle-login' },
      { path: '/logout', action: this.logout },
      { path: '/user-applications', component: 'noodle-user-applications' },
      {
        path: '(.*)',
        redirect: '/dash',
      },
    ]);
    if (!NoodleWeb.IsAuthenticated()) {
      Router.go(`/login`);
    }
  }

  render() {
    return html`<main></main>`;
  }
}
