package com.idoceb00.elogap.http.controller;

import com.idoceb00.elogap.domain.Activity;
import com.idoceb00.elogap.domain.MatchResult;
import com.idoceb00.elogap.service.ActivityService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/v1/activities")
public class ActivityController {

    private final ActivityService service;

    public ActivityController(ActivityService service) {
        this.service = service;
    }

    @GetMapping
    public List<Activity> list(
            @RequestParam Optional<MatchResult> result,
            @RequestParam Optional<String> champion
    ){
        return service.list(result, champion);
    }
}
