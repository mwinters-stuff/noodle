import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_form_builder/flutter_form_builder.dart';
import 'package:flutter_svg/svg.dart';
import 'package:form_builder_validators/form_builder_validators.dart';
import 'package:noodlewebclient/repositories/reositories.dart';
import 'package:noodlewebclient/widgets/widgets.dart';

class LoginScreen extends StatefulWidget {
  static const routeName = '/login';

  const LoginScreen({Key? key}) : super(key: key);
  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final _formKey = GlobalKey<FormBuilderState>();
  String _errorText = "";

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: FormBuilder(
        key: _formKey,
        autovalidateMode: AutovalidateMode.disabled,
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              SvgPicture.asset(
                "images/noodle-icon.svg",
                width: 256,
                height: 256,
              ),
              Text(
                "Noodle",
                textAlign: TextAlign.center,
                style: Theme.of(context).textTheme.displayLarge,
              ),
              const SizedBox(height: 10),
              SizedBox(
                width: 300,
                child: FormBuilderTextField(
                  expands: false,
                  autovalidateMode: AutovalidateMode.onUserInteraction,
                  name: 'username',
                  decoration: const InputDecoration(
                    labelText: 'Username',
                  ),
                  validator: FormBuilderValidators.compose([
                    FormBuilderValidators.required(),
                  ]),
                  keyboardType: TextInputType.text,
                  textInputAction: TextInputAction.next,
                ),
              ),
              const SizedBox(height: 10),
              SizedBox(
                width: 300,
                child: FormBuilderTextField(
                  autovalidateMode: AutovalidateMode.onUserInteraction,
                  name: 'password',
                  expands: false,
                  obscureText: true,
                  decoration: const InputDecoration(
                    labelText: 'Password',
                  ),
                  validator: FormBuilderValidators.required(),
                  keyboardType: TextInputType.text,
                  textInputAction: TextInputAction.next,
                ),
              ),
              const SizedBox(height: 10),
              FilledButton(
                onPressed: () => _login(context),
                child: const Text("Login"),
              ),
              const SizedBox(height: 10),
              Text(
                _errorText,
                style: Theme.of(context).textTheme.bodyMedium!.copyWith(color: Colors.red),
              )
            ],
          ),
        ),
      ),
    );
  }

  void _login(BuildContext context) {
    setState(() => _errorText = '');

    if (_formKey.currentState?.saveAndValidate() ?? false) {
      final username = _formKey.currentState!.fields['username']!.value;
      final password = _formKey.currentState!.fields['password']!.value;
      RepositoryProvider.of<SwaggerRepository>(context)
          .login(username, password)
          .then(
            (success) => {
              if (success) {Navigator.of(context).pushNamed(MainScreen.routeName)} else {setState(() => _errorText = 'Login Failed')}
            },
          )
          .catchError((error) {
        setState(() => _errorText = error.toString().replaceAll("Exception:", ""));
      });
    }
  }
}
