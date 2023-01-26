// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'noodle_service.swagger.dart';

// **************************************************************************
// ChopperGenerator
// **************************************************************************

// ignore_for_file: always_put_control_body_on_new_line, always_specify_types, prefer_const_declarations, unnecessary_brace_in_string_interps
class _$NoodleService extends NoodleService {
  _$NoodleService([ChopperClient? client]) {
    if (client == null) return;
    this.client = client;
  }

  @override
  final definitionType = NoodleService;

  @override
  Future<Response<Object>> _healthzGet() {
    final Uri $url = Uri.parse('/healthz');
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
    );
    return client.send<Object, Object>($request);
  }

  @override
  Future<Response<Object>> _readyzGet() {
    final Uri $url = Uri.parse('/readyz');
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
    );
    return client.send<Object, Object>($request);
  }

  @override
  Future<Response<List<User>>> _noodleUsersGet({
    int? userid,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/users');
    final Map<String, dynamic> $params = <String, dynamic>{'userid': userid};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<User>, User>($request);
  }

  @override
  Future<Response<List<Group>>> _noodleGroupsGet({
    int? groupid,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/groups');
    final Map<String, dynamic> $params = <String, dynamic>{'groupid': groupid};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<Group>, Group>($request);
  }

  @override
  Future<Response<List<UserGroup>>> _noodleUserGroupsGet({
    int? userid,
    int? groupid,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/user-groups');
    final Map<String, dynamic> $params = <String, dynamic>{
      'userid': userid,
      'groupid': groupid,
    };
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<UserGroup>, UserGroup>($request);
  }

  @override
  Future<Response<dynamic>> _noodleLdapReloadGet({String? remoteUser}) {
    final Uri $url = Uri.parse('/noodle/ldap/reload');
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }

  @override
  Future<Response<dynamic>> _noodleHeimdallReloadGet({String? remoteUser}) {
    final Uri $url = Uri.parse('/noodle/heimdall/reload');
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }

  @override
  Future<Response<List<Tab>>> _noodleTabsGet({String? remoteUser}) {
    final Uri $url = Uri.parse('/noodle/tabs');
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      headers: $headers,
    );
    return client.send<List<Tab>, Tab>($request);
  }

  @override
  Future<Response<Tab>> _noodleTabsPost({
    required String? action,
    required Tab? tab,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/tabs');
    final Map<String, dynamic> $params = <String, dynamic>{'action': action};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final $body = tab;
    final Request $request = Request(
      'POST',
      $url,
      client.baseUrl,
      body: $body,
      parameters: $params,
      headers: $headers,
    );
    return client.send<Tab, Tab>($request);
  }

  @override
  Future<Response<dynamic>> _noodleTabsDelete({
    required int? tabid,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/tabs');
    final Map<String, dynamic> $params = <String, dynamic>{'tabid': tabid};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'DELETE',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }

  @override
  Future<Response<List<ApplicationTab>>> _noodleApplicationTabsGet({
    required int? tabId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/application-tabs');
    final Map<String, dynamic> $params = <String, dynamic>{'tab_id': tabId};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<ApplicationTab>, ApplicationTab>($request);
  }

  @override
  Future<Response<ApplicationTab>> _noodleApplicationTabsPost({
    required String? action,
    required ApplicationTab? applicationTab,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/application-tabs');
    final Map<String, dynamic> $params = <String, dynamic>{'action': action};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final $body = applicationTab;
    final Request $request = Request(
      'POST',
      $url,
      client.baseUrl,
      body: $body,
      parameters: $params,
      headers: $headers,
    );
    return client.send<ApplicationTab, ApplicationTab>($request);
  }

  @override
  Future<Response<dynamic>> _noodleApplicationTabsDelete({
    required int? applicationTabId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/application-tabs');
    final Map<String, dynamic> $params = <String, dynamic>{
      'application_tab_id': applicationTabId
    };
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'DELETE',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }

  @override
  Future<Response<List<UserApplications>>> _noodleUserApplicationsGet({
    required int? userId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/user-applications');
    final Map<String, dynamic> $params = <String, dynamic>{'user_id': userId};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<UserApplications>, UserApplications>($request);
  }

  @override
  Future<Response<UserApplications>> _noodleUserApplicationsPost({
    required UserApplications? userApplication,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/user-applications');
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final $body = userApplication;
    final Request $request = Request(
      'POST',
      $url,
      client.baseUrl,
      body: $body,
      headers: $headers,
    );
    return client.send<UserApplications, UserApplications>($request);
  }

  @override
  Future<Response<dynamic>> _noodleUserApplicationsDelete({
    required int? userApplicationId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/user-applications');
    final Map<String, dynamic> $params = <String, dynamic>{
      'user_application_id': userApplicationId
    };
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'DELETE',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }

  @override
  Future<Response<List<GroupApplications>>> _noodleGroupApplicationsGet({
    required int? groupId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/group-applications');
    final Map<String, dynamic> $params = <String, dynamic>{'group_id': groupId};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<GroupApplications>, GroupApplications>($request);
  }

  @override
  Future<Response<GroupApplications>> _noodleGroupApplicationsPost({
    required GroupApplications? groupApplication,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/group-applications');
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final $body = groupApplication;
    final Request $request = Request(
      'POST',
      $url,
      client.baseUrl,
      body: $body,
      headers: $headers,
    );
    return client.send<GroupApplications, GroupApplications>($request);
  }

  @override
  Future<Response<dynamic>> _noodleGroupApplicationsDelete({
    required int? groupApplicationId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/group-applications');
    final Map<String, dynamic> $params = <String, dynamic>{
      'group_application_id': groupApplicationId
    };
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'DELETE',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }

  @override
  Future<Response<List<ApplicationTemplate>>> _noodleAppTemplatesGet({
    required String? search,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/app-templates');
    final Map<String, dynamic> $params = <String, dynamic>{'search': search};
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client
        .send<List<ApplicationTemplate>, ApplicationTemplate>($request);
  }

  @override
  Future<Response<List<Application>>> _noodleApplicationsGet({
    int? applicationId,
    String? applicationTemplate,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/applications');
    final Map<String, dynamic> $params = <String, dynamic>{
      'application_id': applicationId,
      'application_template': applicationTemplate,
    };
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'GET',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<List<Application>, Application>($request);
  }

  @override
  Future<Response<Application>> _noodleApplicationsPost({
    required Application? application,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/applications');
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final $body = application;
    final Request $request = Request(
      'POST',
      $url,
      client.baseUrl,
      body: $body,
      headers: $headers,
    );
    return client.send<Application, Application>($request);
  }

  @override
  Future<Response<dynamic>> _noodleApplicationsDelete({
    required int? applicationId,
    String? remoteUser,
  }) {
    final Uri $url = Uri.parse('/noodle/applications');
    final Map<String, dynamic> $params = <String, dynamic>{
      'application_id': applicationId
    };
    final Map<String, String> $headers = {
      if (remoteUser != null) 'Remote-User': remoteUser,
    };
    final Request $request = Request(
      'DELETE',
      $url,
      client.baseUrl,
      parameters: $params,
      headers: $headers,
    );
    return client.send<dynamic, dynamic>($request);
  }
}
