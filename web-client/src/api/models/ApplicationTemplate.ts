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
/**
 *
 * @export
 * @interface ApplicationTemplate
 */
export interface ApplicationTemplate {
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  Appid?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  Name?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  Website?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  License?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  Description?: string;
  /**
   *
   * @type {boolean}
   * @memberof ApplicationTemplate
   */
  Enhanced?: boolean;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  TileBackground?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  Icon?: string;
  /**
   *
   * @type {string}
   * @memberof ApplicationTemplate
   */
  SHA?: string;
}

/**
 * Check if a given object implements the ApplicationTemplate interface.
 */
export function instanceOfApplicationTemplate(value: object): boolean {
  let isInstance = true;

  return isInstance;
}

export function ApplicationTemplateFromJSON(json: any): ApplicationTemplate {
  return ApplicationTemplateFromJSONTyped(json, false);
}

export function ApplicationTemplateFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean
): ApplicationTemplate {
  if (json === undefined || json === null) {
    return json;
  }
  return {
    Appid: !exists(json, 'Appid') ? undefined : json['Appid'],
    Name: !exists(json, 'Name') ? undefined : json['Name'],
    Website: !exists(json, 'Website') ? undefined : json['Website'],
    License: !exists(json, 'License') ? undefined : json['License'],
    Description: !exists(json, 'Description') ? undefined : json['Description'],
    Enhanced: !exists(json, 'Enhanced') ? undefined : json['Enhanced'],
    TileBackground: !exists(json, 'tile_background')
      ? undefined
      : json['tile_background'],
    Icon: !exists(json, 'Icon') ? undefined : json['Icon'],
    SHA: !exists(json, 'SHA') ? undefined : json['SHA'],
  };
}

export function ApplicationTemplateToJSON(
  value?: ApplicationTemplate | null
): any {
  if (value === undefined) {
    return undefined;
  }
  if (value === null) {
    return null;
  }
  return {
    Appid: value.Appid,
    Name: value.Name,
    Website: value.Website,
    License: value.License,
    Description: value.Description,
    Enhanced: value.Enhanced,
    tile_background: value.TileBackground,
    Icon: value.Icon,
    SHA: value.SHA,
  };
}
