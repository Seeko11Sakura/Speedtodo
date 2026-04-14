import 'package:dida_app/features/tasks/presentation/task_list_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('renders add task button', (tester) async {
    await tester.pumpWidget(const MaterialApp(home: TaskListPage(view: 'inbox')));
    expect(find.text('Add Task'), findsOneWidget);
  });
}
