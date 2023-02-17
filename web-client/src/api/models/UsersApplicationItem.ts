/* tslint:disable */
/* eslint-disable */
/**
 * Noodle
 * Noodle
 *
 * The version of the OpenAPI document: 2.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { Application } from './Application';
import {
  ApplicationFromJSON,
  ApplicationFromJSONTyped,
  ApplicationToJSON,
} from './Application';

/**
 *
 * @export
 * @interface UsersApplicationItem
 */
export interface UsersApplicationItem {
  /**
   *
   * @type {number}
   * @memberof UsersApplicationItem
   */
  tabId?: number;
  /**
   *
   * @type {number}
   * @memberof UsersApplicationItem
   */
  displayOrder?: number;
  /**
   *
   * @type {Application}
   * @memberof UsersApplicationItem
   */
  application?: Application;
}

/**
 * Check if a given object implements the UsersApplicationItem interface.
 */
export function instanceOfUsersApplicationItem(value: object): boolean {
  let isInstance = true;

  return isInstance;
}

export function UsersApplicationItemFromJSON(json: any): UsersApplicationItem {
  return UsersApplicationItemFromJSONTyped(json, false);
}

export function UsersApplicationItemFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): UsersApplicationItem {
  if (json === undefined || json === null) {
    return json;
  }
  return {
    tabId: !exists(json, 'TabId') ? undefined : json['TabId'],
    displayOrder: !exists(json, 'DisplayOrder')
      ? undefined
      : json['DisplayOrder'],
    application: !exists(json, 'Application')
      ? undefined
      : ApplicationFromJSON(json['Application']),
  };
}

export function UsersApplicationItemToJSON(
  value?: UsersApplicationItem | null
): any {
  if (value === undefined) {
    return undefined;
  }
  if (value === null) {
    return null;
  }
  return {
    TabId: value.tabId,
    DisplayOrder: value.displayOrder,
    Application: ApplicationToJSON(value.application),
  };
}