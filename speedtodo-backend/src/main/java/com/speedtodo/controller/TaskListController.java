package com.speedtodo.controller;
import com.speedtodo.model.TaskList;
import com.speedtodo.repository.TaskListRepository;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.net.URI;
import java.util.List;
@RestController
@RequestMapping("/lists")
public class TaskListController {
    private final TaskListRepository repository;
    public TaskListController(TaskListRepository repository) { this.repository = repository; }
    @GetMapping public List<TaskList> list() { return repository.findAll(); }
    @PostMapping public ResponseEntity<TaskList> create(@RequestBody TaskList list) { TaskList created = repository.save(list); return ResponseEntity.created(URI.create("/lists/" + created.getId())).body(created); }
    @PatchMapping("/{id}") public TaskList update(@PathVariable Long id, @RequestBody TaskList list) {
        TaskList existing = repository.findById(id).orElseThrow(() -> new RuntimeException("Not found"));
        if (list.getName() != null) existing.setName(list.getName());
        if (list.getColor() != null) existing.setColor(list.getColor());
        existing.setSortOrder(list.getSortOrder());
        existing.setDefault(list.isDefault());
        return repository.save(existing);
    }
}