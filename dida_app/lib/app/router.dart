import 'package:go_router/go_router.dart';

import '../features/tasks/presentation/task_list_page.dart';
import '../features/views/presentation/today_page.dart';
import '../features/views/presentation/upcoming_page.dart';
import '../features/views/presentation/calendar_page.dart';
import '../features/pomodoro/presentation/pomodoro_page.dart';
import '../features/lists/presentation/list_page.dart';
import '../features/tags/presentation/tag_page.dart';

final appRouter = GoRouter(
  initialLocation: '/today',
  routes: [
    GoRoute(path: '/inbox', builder: (_, __) => const TaskListPage(view: 'inbox')),
    GoRoute(path: '/today', builder: (_, __) => const TodayPage()),
    GoRoute(path: '/upcoming', builder: (_, __) => const UpcomingPage()),
    GoRoute(path: '/calendar', builder: (_, __) => const CalendarPage()),
    GoRoute(path: '/pomodoro', builder: (_, __) => const PomodoroPage()),
    GoRoute(path: '/lists', builder: (_, __) => const ListPage()),
    GoRoute(path: '/tags', builder: (_, __) => const TagPage()),
  ],
);
