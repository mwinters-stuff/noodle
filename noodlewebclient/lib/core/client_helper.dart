import 'package:noodlewebclient/main.dart';
import 'package:noodlewebclient/swagger_generated/noodle_service.swagger.dart';

class NoodleClientHelper {
  static NoodleService? _noodleClient;

  static NoodleService getClient() {
    if (_noodleClient == null) {
      throw Exception('Client is not yet created');
    }
    //It can't be null here, to compile bang operator is needed
    return _noodleClient!;
  }

  static NoodleService initClient(String serverUrl, {String? username, String? password}) {
    _noodleClient = NoodleService.create(baseUrl: Uri.parse(serverUrl));
    return getClient();
  }

  static bool initNeeded() => _noodleClient == null;

  static Future clearAuthStorage() async {
    await preferences.remove('username');
    await preferences.remove('password');
  }
}
