import '../swagger_generated/noodle_service.swagger.dart';

class SwaggerRepository {
  final NoodleService noodleService;
  String token = "";

  SwaggerRepository(String serverUrl) : noodleService = NoodleService.create(baseUrl: Uri.parse(serverUrl));

  Future<bool> login(String username, String password) async {
    var usersession = await noodleService.authAuthenticatePost(
      login: UserLogin(
        username: username,
        password: password,
      ),
    );

    if (usersession.isSuccessful) {
      token = usersession.body?.token ?? "";
    } else {
      return Future.error(Exception(usersession.base.reasonPhrase));
    }

    return usersession.isSuccessful && token != "";
  }
}
