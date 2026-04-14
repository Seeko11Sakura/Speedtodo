import 'package:dida_app/features/pomodoro/presentation/pomodoro_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('shows start focus action', (tester) async {
    await tester.pumpWidget(const MaterialApp(home: PomodoroPage()));
    expect(find.text('Start Focus'), findsOneWidget);
  });
}
