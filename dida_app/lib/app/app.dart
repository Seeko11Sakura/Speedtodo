import 'package:flutter/material.dart';

import 'router.dart';
import '../core/theme/app_theme.dart';

class DidaApp extends StatelessWidget {
  const DidaApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Dida MVP',
      theme: buildAppTheme(),
      routerConfig: appRouter,
    );
  }
}
