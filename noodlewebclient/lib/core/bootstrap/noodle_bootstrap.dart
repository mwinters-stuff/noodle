import 'package:flutter/material.dart';
import 'package:noodlewebclient/core/bootstrap/bootSteps/login_boot_sttep.dart';
import 'package:noodlewebclient/core/bootstrap/bootstrap.dart';

class NoodleBootstrap extends StatelessWidget {
  final BootCompleted bootCompleted;

  const NoodleBootstrap(this.bootCompleted, {Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BootStrap(
      [
        LoginBootStep(),
      ],
      bootCompleted,
    );
  }
}
/*
class CaladriusBootstrapCustom extends StatelessWidget {
  final BootCompleted bootCompleted;
  final loginBootStep = const LoginBootStep();

  const CaladriusBootstrapCustom(
    this.bootCompleted, {
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    if (loginBootStep.stepRequired()) {
      return BootStrap(
        [
          loginBootStep,
        ],
      );
    }
    return bootCompleted();
  }
}
*/
