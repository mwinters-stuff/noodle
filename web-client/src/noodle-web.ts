/* eslint-disable no-console */
import { html, css, LitElement } from 'lit';
import { Commands, Context, Router } from '@vaadin/router';
import { customElement, state } from 'lit/decorators.js';

import './noodle-login.js';
import './dash/noodle-dash.js';
import './user-applications/noodle-user-applications.js';
import './settings/noodle-settings.js';

import { setBasePath } from '@shoelace-style/shoelace/dist/utilities/base-path.js';


import { provide } from '@lit-labs/context';
import {
  Application,
  Configuration,
  ConfigurationParameters,
  NoodleApiApi,
  NoodleAuthApi,
  Tab,
  UserSession,
  UsersApplicationItem,
} from './api/index.js';
import {
  authApiContext,
  dataCacheContext,
  noodleApiContext,
  userSessionContext,
  DataCache,
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

  @provide({ context: dataCacheContext })
  @state()
  dataCache: DataCache = {
    _applications: [],
    _tabs: [],
    router: undefined,
    _userApplications: [],

    SetUserApplications(uai: UsersApplicationItem[]): void {
      this._userApplications = uai;
      this._applications = [];
      uai.forEach(userApp => {
        this._applications.push(userApp.Application!);
      });
    },

    GetUserApplicationsForTab(tabId: number): UsersApplicationItem[] {
      return this._userApplications.filter((value) => value.TabId === tabId);
    },

    Applications(): Application[] {
      return this._applications;
    },
    Tabs(): Tab[] {
      return this._tabs;
    },
    GetApplication(id: number): Application {
      return this.Applications().find(value => value.Id === id)!;
    },
    UserApplications(): UsersApplicationItem[] {
      return this._userApplications;
    },


    SetTabs(value: Tab[]): void {
      this._tabs = value;
    },
    GetTabIndex(tabId: number): number {
      return this._tabs.findIndex(tab => tab.Id === tabId);
    },
  };

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

  @state()
  activeRoute: string = '';

  @state()
  routeParams: any = {};

  @state()
  routeQuery: Object = {};

  @state()
  routeData: Object = {};

  constructor() {
    super();
    const parameters: ConfigurationParameters = {
      apiKey: NoodleWeb.apiKey,
    };
    const config = new Configuration(parameters);

    this.noodleApi = new NoodleApiApi(config);
    this.authApi = new NoodleAuthApi(config);

    const token = NoodleWeb.getAuthToken();

    // Set the base path to the folder you copied Shoelace's assets to
    setBasePath('node_modules/@shoelace-style/shoelace/dist');

    if (token !== '') {
      this.authApi
        .authSessionGet(token)
        .then(value => {
          this.userSession = value;
          this.RefreshTabs();
        })
        .catch(reason => {
          console.log('GetUserSession Error: ', reason);
        });
    }
  }

  private static apiKey(name: string): string {
    console.log(`GetAPIKey ${name}`);
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

  private RefreshTabs() {
    this.noodleApi.noodleTabsGet().then(value => {
      this.dataCache.SetTabs(value);
    });
  }

  firstUpdated() {
    const router = new Router(this.shadowRoot!.querySelector('main'));
    router.setRoutes([
      { path: '/', component: 'noodle-dash' },
      { path: '/dash/:tabId', component: 'noodle-dash' },
      { path: '/login', component: 'noodle-login' },
      { path: '/logout', action: this.logout },
      { path: '/user-applications', component: 'noodle-user-applications' },
      { path: '/settings', component: 'noodle-settings' },
      {
        path: '(.*)',
        redirect: '/dash',
      },
    ]);
    this.dataCache.router = router;
    if (!NoodleWeb.IsAuthenticated()) {
      Router.go(`/login`);
    }
  }

  render() {
    return html`<main></main>`;
  }
}
