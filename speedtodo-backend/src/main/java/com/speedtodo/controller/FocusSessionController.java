package com.speedtodo.controller;
import com.speedtodo.model.FocusSession;
import com.speedtodo.service.FocusSessionService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.net.URI;
import java.util.List;
@RestController
@RequestMapping("/focus-sessions")
public class FocusSessionController {
    private final FocusSessionService service;
    public FocusSessionController(FocusSessionService service) { this.service = service; }
    @GetMapping public List<FocusSession> list() { return service.list(); }
    @PostMapping public ResponseEntity<FocusSession> create(@RequestBody FocusSession session) { FocusSession created = service.create(session); return ResponseEntity.created(URI.create("/focus-sessions/" + created.getId())).body(created); }
    @PostMapping("/{id}/complete") public FocusSession complete(@PathVariable Long id) { return service.complete(id); }
    @PostMapping("/{id}/cancel") public FocusSession cancel(@PathVariable Long id) { return service.cancel(id); }
}