Feature: change order status
    Scenario: change order to received
        When I change from pending to received
        Then the result should equal true

    Scenario: change order to in prepare
        When I change from received to in_prepare
        Then the result should equal true

    Scenario: change order to done
        When I change from in_prepare to done
        Then the result should equal true

    Scenario: change order to finished
        When I change from done to finished
        Then the result should equal true

    Scenario: change order to invalid status
        When I change from finished to received
        Then the result should equal false
