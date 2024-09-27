Find the last time each bike was in use. Output both the bike number and the date-timestamp of the 
bike's last use (i.e., the date-time the bike was returned). Order the results by bikes that were most recently used.

so I solved it i had visualisation

here what i visualised 

- what it is asking
    - hume vo bike numbers do jo last used hue hai

    - and data set is looking like 
        - it has multiple records of bike usage which has 
            - bike pickup starting time (start_time) and bike returning time (end_time)

            - to hume kaise pta chale ki last time bike kab use hui thi sql data dekh ke
                - hume ye find krna hai ki last usage record konsa hai
                    - vo aisa hoga ki hum end_time ka max aggregate nikalenge 
                    - or AGGREGATE hamesha GROUP BY ke saath use hota hai
                        - to ab hume har bike ka last use (max of end_time) chaiye 
                            - to hum uss poore data set ke common bike_number records ko individual bike_number me merge (sql me GROUP BY) kr denge GROUP BY use krke
                            - to hume bike_number column me unique records mil jaenge common merge ho jaenge or kyuki humne AGGREGATE use krke 

Solution - 

SELECT bike_number,
       max(end_time) last_used
FROM dc_bikeshare_q1_2012
GROUP BY bike_number
ORDER BY last_used DESC

